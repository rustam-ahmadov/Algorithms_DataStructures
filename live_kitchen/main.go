package main

import (
	"context"
	"time"
)

//Сервис кухни для обработки заказов в реальном времени
//Сервис забирает заказ из очереди и пишет в очередь когда заказ готов. После берет следующий.
//На кухне может быть несколько поворов, между которыми распределены:

//1 Обработка заказа из очереди может занимать 2 минуты
//2 Таймаут на чтение из кафки 10 сек
//3 Когда повар приготовил, он жмет свою кнопку, Сервис кухни пишет пишет в
//		кафку что заказ готов и берет новый заказ из очереди для освободившегося поворя
//4 Если заказ отменили то повар должен прекратить заниматься заказом и взять новый заказ

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Msg struct {
	ID     int32  `json:"id"`
	CookID int32  `json:"cook_id"`
	Status string `json:"status"`
}

func (s *Service) Process(ctx context.Context, msg *Msg) error {
	if err := s.process(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (s *Service) process(ctx context.Context, msg *Msg) error {
	time.Sleep(time.Minute * 2)

	return nil
}

func main() {

}
