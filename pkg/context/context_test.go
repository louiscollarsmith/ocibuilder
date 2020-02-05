/*
Copyright 2019 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ocibuilder/ocibuilder/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestInjectDockerfile(t *testing.T) {
	err := ioutil.WriteFile("../../testing/dummy/Dockerfile", []byte("FROM alpine\nCOPY . .\n"), 0644)
	assert.Equal(t, nil, err)

	err = util.TarFile([]string{"../../testing/e2e/resources/go-test-service"}, "../../testing/e2e/resources/go-test-service/ocib/context/context.tar.gz")
	assert.Equal(t, nil, err)

	err = InjectDockerfile("../../testing/e2e/resources/go-test-service", "../../testing/dummy/Dockerfile")
	assert.Equal(t, nil, err)

	err = os.RemoveAll("../../testing/e2e/resources/go-test-service/ocib")
	assert.Equal(t, nil, err)
}

func TestExcludeIgnored(t *testing.T) {
	const DIR = "../../testing/e2e/resources/go-test-service/"

	files, err := ExcludeIgnored("../../testing/e2e/resources/go-test-service")
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{DIR + ".dockerignore", DIR + "main.go", DIR + "ocibuilder.yaml", DIR + "overlay.yaml"}, files)
}
