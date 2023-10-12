/*
 * @Date: 2023-10-12 00:58:01
 * @LastEditors: HeXu
 * @LastEditTime: 2023-10-12 00:58:35
 * @FilePath: /crawler/engine/shedule.go
 */
package engine

import "github.com/heexu976/crawler/collect"

type ScheduleEngine struct {
	requestCh chan *collect.Request
}