package rest

func (h *DeliveryHttpEngine) initRoute() error {
	h.mux.Get("/", IndexHandler)

	return nil
}
