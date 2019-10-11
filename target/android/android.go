package android

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/qlova/seed"
)

//TODO Use https://developer.android.com/reference/androidx/webkit/WebViewAssetLoader
//Needs https://android.googlesource.com/platform/frameworks/support/+/androidx-master-dev/webkit/src/main/java/androidx/webkit/WebViewAssetLoader.java
//And https://android.googlesource.com/platform/frameworks/support/+/androidx-master-dev/webkit/src/main/java/androidx/webkit/internal/AssetHelper.java
const MainActivity = `package com.example.app;

import android.widget.RelativeLayout;
import android.widget.FrameLayout;
import android.annotation.SuppressLint;
import android.app.Activity;
import android.os.Bundle;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.os.Build;

public class MainActivity extends Activity {

    private WebView w;

    @SuppressLint("SetJavaScriptEnabled")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
		
		w = new WebView(this);
		RelativeLayout.LayoutParams params = new RelativeLayout.LayoutParams(RelativeLayout.LayoutParams.FILL_PARENT, 100);
		params.addRule(RelativeLayout.ALIGN_PARENT_BOTTOM);
		w.setLayoutParams(params);
		
		setContentView(w);

        // Force links and redirects to open in the WebView instead of in a browser
        w.setWebViewClient(new WebViewClient());

        // Enable Javascript
        WebSettings webSettings = w.getSettings();
		webSettings.setJavaScriptEnabled(true);
		
		if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT) {
			w.setWebContentsDebuggingEnabled(true);
		}

		w.getSettings().setDomStorageEnabled(true);

        w.loadUrl("file:///android_asset/index.html");
    }

    // Prevent the back-button from closing the app
    @Override
    public void onBackPressed() {
        if(w.canGoBack()) {
            w.goBack();
        } else {
            super.onBackPressed();
        }
    }
}`

func Launch(app *seed.App) error {
	var base = app.Name
	var SDK = Home() + "/.qlovaseed/android"
	var packagename = "com.example.app"

	err := Export(app)
	if err != nil {
		return err
	}

	cmd := exec.Command(SDK+"/adb", "install", "-r", base+".apk")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	var args = []string{"shell", "am", "start", "-n", packagename + "/" + packagename + ".MainActivity"}

	cmd = exec.Command(SDK+"/adb", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}

//Export creates an apk of the app.
func Export(app *seed.App) error {
	var base = app.Name

	var SDK = Home() + "/.qlovaseed/android"
	if _, err := os.Stat(SDK + "/sdk.jar"); os.IsNotExist(err) {
		fmt.Println("Downloading the android sdk... (15mb~)")
		os.MkdirAll(SDK, 0755)
		err := DownloadFile(SDK+"/../android.zip", "https://bitbucket.org/Splizard/ilang-release/downloads/android.zip")
		if err != nil {
			return err
		}
		err = Unzip(SDK+"/../android.zip", SDK)
		if err != nil {
			return err
		}
	}

	var temp, err = ioutil.TempDir("", "qlovaseed")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(temp+"/MainActivity.java", []byte(MainActivity), 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(temp+"/AndroidManifest.xml", []byte(`<?xml version="1.0" encoding="utf-8"?>
	<manifest xmlns:android="http://schemas.android.com/apk/res/android"
		package="com.example.app">
		
		<uses-permission android:name="android.permission.INTERNET" />
		<uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
		<uses-permission android:name="android.permission.READ_PHONE_STATE" />
		<uses-permission android:name="android.permission.ACCESS_WIFI_STATE" />

		<uses-sdk
			android:minSdkVersion="21"
			android:targetSdkVersion="29" />

		<application
			android:allowBackup="true"
			android:icon="@drawable/ic_launcher"
			android:label="Example"
			android:hardwareAccelerated="true"
			android:theme="@android:style/Theme.NoTitleBar">
			
			<activity android:name=".MainActivity">
				<intent-filter>
					<action android:name="android.intent.action.MAIN" />
					<category android:name="android.intent.category.LAUNCHER" />
				</intent-filter>
			</activity>
		</application>
	</manifest>`), 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("javac", "-classpath", SDK+"/sdk.jar", "-sourcepath", "src:gen", "-d", ".",
		"-target", "1.7", "-source", "1.7", "MainActivity.java")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command(SDK+"/dx", "--dex", "--output=classes.dex", ".")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	//Move resources into the correct places.
	Unzip(SDK+"/res.zip", temp+"/res")

	cmd = exec.Command(SDK+"/aapt", "package", "-f", "-M", "AndroidManifest.xml", "-S", "res", "-I", SDK+"/sdk.jar", "-F", base+".apk.unaligned")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := os.MkdirAll(seed.Dir+"/assets", 0755); err != nil {
		return err
	}

	//Add assets.
	var data = seed.Dir + "/assets"
	linkto, err := filepath.Abs(data)
	if err != nil {
		return err
	}
	os.Symlink(linkto, temp+"/assets")

	index, err := os.Create(temp + "/assets/index.html")
	if err != nil {
		return err
	}

	index.Write(app.Render(seed.Default))
	index.Close()

	cmd = exec.Command(SDK+"/aapt", "add", base+".apk.unaligned", "assets/index.html")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return err
	}

	for name, data := range seed.GetEmbeddings() {
		var file, err = os.Create(temp + "/assets" + name)
		if err != nil {
			return err
		}
		file.Write(data.Data)
		file.Close()

		cmd = exec.Command(SDK+"/aapt", "add", base+".apk.unaligned", "assets"+name)
		cmd.Dir = temp
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			return err
		}
	}

	/*var FilePathWalkDir = func(root string) ([]string, error) {
		var files []string
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			path = strings.Replace(path, data, "assets/", 1)
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		return files, err
	}

	if _, err := os.Stat(data); err == nil {
		files, err := FilePathWalkDir(data)
		if err != nil {
			return err
		}

		for _, file := range files {
			cmd = exec.Command(SDK+"/aapt", "add", base+".apk.unaligned", file)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				return err
			}
		}
	}*/

	cmd = exec.Command(SDK+"/aapt", "add", base+".apk.unaligned", "classes.dex")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("jarsigner", "-keystore", SDK+"/debug.keystore", "-storepass", "android", base+".apk.unaligned", "androiddebugkey")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command(SDK+"/zipalign", "-f", "4", base+".apk.unaligned", base+"-debug.apk")
	cmd.Dir = temp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := CopyFile(temp+"/"+base+"-debug.apk", seed.Dir+"/"+base+".apk", ""); err != nil {
		return err
	}

	if err := os.RemoveAll(temp); err != nil {
		return err
	}

	return nil
}

func Home() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}
	return ""
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst, preface string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	err = copyFileContents(src, dst, preface)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst, preface string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if preface != "" {
		out.Write([]byte(preface))
		out.Sync()
	}
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
