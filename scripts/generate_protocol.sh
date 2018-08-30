#!/bin/bash

cur_dir=$(cd `dirname $0` && pwd)
root_dir=$(cd $cur_dir/.. && pwd)
source_dir=$root_dir/api/protobuf
output_dir=$root_dir/pkg/protocol/pb

rm -rf $output_dir/*

protoc -I=$source_dir --go_out=$output_dir $source_dir/dvc.proto