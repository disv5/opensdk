/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 否定词重复而导致失败的否定词列表，包括请求的否定词之间重复、请求的否定词和已有否定词重复、否定词和广告组中关键词重复等
type DuplicateNegativeWordStruct struct {
	PhraseNegativeWords *[]string `json:"phrase_negative_words,omitempty"`
	ExactNegativeWords  *[]string `json:"exact_negative_words,omitempty"`
}
