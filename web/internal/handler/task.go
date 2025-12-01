func (h *Handler) TaskList(c *fiber.Ctx) error {
	tasks, _ := h.service.GetTasks(c.Context())
	return components.TaskList(tasks).Render(c.Context(), c.Response().BodyWriter())
}
