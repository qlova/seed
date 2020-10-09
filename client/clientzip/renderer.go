package clientzip

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
)

func init() {
	client.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
			window.clientzip = {};

			clientzip.FilesAndDirectories = async function(input, policy) {
				let count = 0;
				let single = null;
				let name = "untitled";

				let zip = new JSZip();

				let process = async function(entries, path) {
					for (let entry of entries) {
						if (count == 0) {
							name = entry.name
						}
						count++;

						if ("getFilesAndDirectories" in entry) {
							process(await entry.getFilesAndDirectories(), entry.path + "/");
						} else {
							if (entry.name) {
								zip.file((path + entry.name).replace(/^[/\\]/, ""), entry);
								if (count == 1) single = entry;
							}
						}
					}
				};
				await process(await input.getFilesAndDirectories(), "");

				if (single && count == 1 && policy == "folders") {
					return new File([single], name);
				}

				return new File([await zip.generateAsync({type: "blob"})], name+".zip");
			};

			clientzip.WebkitRelativePaths = async function(input, policy) {
				let count = 0;
				let single = null;
				let name = "untitled";

				let zip = new JSZip();

				for (let file of input.files) {
					if (count == 0) {
						name = file.webkitRelativePath.match(/[^/\\]+/)
						if (name != file.name) count++;
					}

					zip.file(file.webkitRelativePath || file.name, file);
					count++;
					if (count == 1) single = file;
				}

				if (single && count == 1 && policy == "folders") {
					return new File([single], name);
				}

				return new File([await zip.generateAsync({type: "blob"})], name+".zip");
			};

			clientzip.AsEntries = async function(items, policy) {
				let count = 0;
				let single = null;
				let name = "untitled";

				async function process(entry) {
					if (count == 0) {
						name = entry.name
					}
					count++;

					if (entry && entry.isFile) {
						let p = await new Promise((resolve, reject) => entry.file(resolve, reject));
						let file = await p;
						zip.file(file.name, file);
						if (count == 1) single = file;
					} else if (entry && entry.isDirectory) {
						let reader = entry.createReader();
						let entries = await new Promise((resolve, reject) => reader.readEntries(resolve, reject));

						for (let entry of entries) {
							await process(entry)
						}
					}
				}

				let zip = new JSZip();

				for (let item of items) {
					await process(item.webkitGetAsEntry());
				}

				if (single && count == 1 && policy == "folders") {
					return new File([single], name);
				}

				return new File([await zip.generateAsync({type: "blob"})], name+".zip");
			};
		`)
	})
}
