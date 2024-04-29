/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package libjvm

import (
	"github.com/Masterminds/semver/v3"
)

var Java9, _ = semver.NewVersion("9")
var Java17, _ = semver.NewVersion("17")
var Java18, _ = semver.NewVersion("18")

func IsBeforeJava9(candidate string) bool {
	v, err := semver.NewVersion(candidate)
	if err != nil {
		return false
	}

	return v.LessThan(Java9)
}

func IsBeforeJava17(candidate string) bool {
	v, err := semver.NewVersion(candidate)
	if err != nil {
		return false
	}

	return v.LessThan(Java17)
}

func IsBeforeJava18(candidate string) bool {
	v, err := semver.NewVersion(candidate)
	if err != nil {
		return false
	}

	return v.LessThan(Java18)
}
