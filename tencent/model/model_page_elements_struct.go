/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 组件素材内容，组件的同步顺序与原生页的展示顺序一致，即第一个同步组件为顶部展示元素
type PageElementsStruct struct {
	ElementShelf        *ElementShelf        `json:"element_shelf,omitempty"`
	ElementFloat        *ElementFloat        `json:"element_float,omitempty"`
	ElementGoods        *ElementGoods        `json:"element_goods,omitempty"`
	ElementSwipe        *ElementSwipe        `json:"element_swipe,omitempty"`
	ElementWebview      *ElementWebview      `json:"element_webview,omitempty"`
	ElementAnimateFloat *ElementAnimateFloat `json:"element_animate_float,omitempty"`
	ImageSpec           *ElementImage        `json:"image_spec,omitempty"`
	VideoSpec           *ElementVideo        `json:"video_spec,omitempty"`
	TextSpec            *ElementText         `json:"text_spec,omitempty"`
	ButtonSpec          *ElementButtonRead   `json:"button_spec,omitempty"`
	FormSpec            *ElementForm         `json:"form_spec,omitempty"`
	ElementType         PageElementsType     `json:"element_type,omitempty"`
}
