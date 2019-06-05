package seed

import "os"
import "path/filepath"

type Target int

const (
	ReactNative Target = iota
	Website
)

func (app App) Export(t Target) {
	if t == Website {
		var dir = filepath.Dir(os.Args[0])

		os.Mkdir(dir+"/website", 0755)

		var index, err = os.Create(dir + "/website/index.html")
		if err != nil {
			panic(err.Error())
		}
		defer index.Close()

		for name, data := range embeddings {
			var file, err = os.Create(dir + "/website/" + name)
			if err != nil {
				panic(err.Error())
			}
			defer file.Close()

			file.Write(data.Data)
		}

		index.Write(app.render(true, Mobile))
		return
	}

	if t == ReactNative {

		var dir = filepath.Dir(os.Args[0])

		os.Mkdir(dir+"/native", 0755)

		var appjs, err = os.Create(dir + "/native/App.js")
		if err != nil {
			panic(err.Error())
		}
		defer appjs.Close()

		appjs.WriteString(`
		import React, { Component } from 'react';
import { LayoutAnimation, TextInput, Text, View, StyleSheet, UIManager } from 'react-native';
import { Constants } from 'expo';

import { WebView } from 'react-native';

const firstHtml =
  '<html><head><style>html, body { margin:0; padding:0; overflow:hidden; background-color: transparent; } svg { position:fixed; top:0; left:0; height:100%; width:100% }</style></head><body>';
const lastHtml = '</body></html>';

class SvgImage extends Component {
  state = { fetchingUrl: null, svgContent: null };
  componentDidMount() {
    this.doFetch(this.props);
  }
  componentWillReceiveProps(nextProps) {
    this.doFetch(nextProps);
  }
  doFetch = props => {
    let uri = props.source && props.source.uri;
    if (uri) {
      props.onLoadStart && props.onLoadStart();
      if (uri.match(/^data:image\/svg/)) {
        const index = uri.indexOf('<svg');
        this.setState({ fetchingUrl: uri, svgContent: uri.slice(index) });
        props.onLoadEnd && props.onLoadEnd();
      } else {
        fetch(uri)
          .then(res => res.text())
          .then(text => {
            this.setState({ fetchingUrl: uri, svgContent: text });
          })
          .catch(err => {
            console.error('got error', err);
          })
          .finally(() => {
            props.onLoadEnd && props.onLoadEnd();
          })
      }
    }
  };
  render() {
    const props = this.props;
    const { svgContent } = this.state;
    if (svgContent) {
      return (
        <View pointerEvents="none" style={[props.style, props.containerStyle]}>
          <WebView
            originWhitelist={['*']}
            scalesPageToFit={true}
            style={[
              {
                width: 200,
                height: 100,
                backgroundColor: 'transparent',
              },
              props.style,
            ]}
            scrollEnabled={false}
            source={{ html: ` + "`" + `${firstHtml}${svgContent}${lastHtml}` + "`" + ` }}
          />
        </View>
      );
    } else {
      return (
        <View
          pointerEvents="none"
          style={[props.containerStyle, props.style]}
        />
      );
    }
  }
}

export default class App extends Component {
	constructor(props) {
		super(props);
		
		UIManager.setLayoutAnimationEnabledExperimental && UIManager.setLayoutAnimationEnabledExperimental(true);
		
		this.state = {
			page: 0,
			
			//All component state is stored here.
			Gg: {
				text: "Hello World",
				style: function() {
					return {
						color: "red"
					}
				}
			}
		};
	}
	
	render() {
	  
	  var global = this;
	  
	var test = function() {		
		LayoutAnimation.easeInEaseOut();
		global.setState(function(old){
			old.Gg.text = "Lel";
			old.Gg.style = function() { 
				return {
					color: "blue"
				}
			}
			old.page = 1;
			return old;
		});
	  }
	  
	return ( <View style={{justifyContent: "center", alignItems: 'center', flex: 1}}>
		`)

		appjs.Write(app.Root().JSX())

		appjs.WriteString(`
	
	</View>
	);
  }
}
		`)

	} else {
		panic("Invalid Export Target")
	}
}
