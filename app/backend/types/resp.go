/*
 * Copyright 2025 Bronya0 <tangssst@163.com>.
 * Author Github: https://github.com/Bronya0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

type Tag struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

type Config struct {
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Language string   `json:"language"`
	Theme    string   `json:"theme"`
	PwdHash  string   `json:"pwd_hash"`
	DBDir    string   `json:"db_dir"`
	DbConf   DBConfig `json:"db_conf" yaml:"db_conf"`
}

type ResultsResp struct {
	Results []any  `json:"results"`
	Err     string `json:"err"`
}
type ResultResp struct {
	Result map[string]any `json:"result"`
	Err    string         `json:"err"`
}

type H map[string]any
