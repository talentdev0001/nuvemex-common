.PHONY: deps test wire install
.SILENT: wire build clean test
SHELL := bash

build_dir := ./.build
package_dir := ${build_dir}
env := ${app_env}

pkg_path="${GOPATH}/pkg/mod/github.com"
goseanto_repo='!montrealist-c!punto/goseanto'
goseanto_version=$(shell cat go.mod | grep -o '/goseanto v[0-9].[0-9].[0-9]' | cut -d' ' -f2)
goseanto_resources="${goseanto_version}/resources/config"

clean:
	rm -rf ${build_dir}

deps:
	go get -u ./...

build: clean
	if [ "${env}" = "" ]; then echo "Please set app_env"; exit 1; fi;
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${build_dir}/search ./lambda/search/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${build_dir}/hinter ./lambda/hinter/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${build_dir}/details ./lambda/details/main.go

	mkdir -p ${build_dir}/config

	# copy goseanto configs
	cp ${pkg_path}/${goseanto_repo}@${goseanto_version}/resources/config/config.yml ${build_dir}/config/goseanto.yml
	cp ${pkg_path}/${goseanto_repo}@${goseanto_version}/resources/config/${env}.yml ${build_dir}/config/goseanto-${env}.yml

	cp ./resources/config/config.yml ${build_dir}/config/
	cp ./resources/config/${env}.yml ${build_dir}/config/${env}.yml

package: build
	@cd ${build_dir} && \
		zip -q -j -r search.zip ./search ./config/ && rm search && \
		zip -q -j -r hinter.zip ./hinter ./config/ && rm hinter && \
		zip -q -j -r details.zip ./details ./config/ && rm details

test:
	go mod download
	cp ${pkg_path}/${goseanto_repo}@${goseanto_version}/resources/config/config.yml ./resources/config/goseanto.yml
	cp ${pkg_path}/${goseanto_repo}@${goseanto_version}/resources/config/testing.yml ./resources/config/goseanto-testing.yml

	go test -v \
			-coverprofile .testCoverage.txt -count=1

	chmod 0777 resources/config/goseanto*

wire:
	wire .

install:
	@./resources/install.sh

docker-build:
	@docker-compose build

docker-up:
	@docker-compose up -d

docker-stop:
	@docker-compose stop

docker-logs:
	@docker-compose logs -f
