package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")

	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-http-frame %v\n", release)
		return
	}

	var flags flag.FlagSet
	mode := flags.String("frame", "gin", "frame: gin, fiber; --frame=gin")
	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f, FrameMode(*mode))
		}
		return nil
	})
}
