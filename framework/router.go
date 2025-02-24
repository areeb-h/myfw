package framework

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
)

// App struct for managing routes
type App struct {
	engine *fiber.App
}

// New initializes the Fiber-powered app
func New() *App {
	return &App{engine: fiber.New()}
}

// Define HTTP methods (Now supports structured return types)
func (app *App) Get(path string, handler any)    { app.engine.Get(path, wrapHandler(handler)) }
func (app *App) Post(path string, handler any)   { app.engine.Post(path, wrapHandlerWithBody(handler)) }
func (app *App) Put(path string, handler any)    { app.engine.Put(path, wrapHandlerWithBody(handler)) }
func (app *App) Patch(path string, handler any)  { app.engine.Patch(path, wrapHandlerWithBody(handler)) }
func (app *App) Delete(path string, handler any) { app.engine.Delete(path, wrapHandlerWithBody(handler)) }

// Start the server
func (app *App) Start() { app.engine.Listen(":8080") }

// Wraps handlers for automatic JSON response handling
func wrapHandler(handler any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := callHandler(handler)
		return c.JSON(response)
	}
}

// Wraps handlers with request body parsing
func wrapHandlerWithBody(handler any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bodyType := getHandlerBodyType(handler)
		if bodyType != nil {
			body := reflect.New(bodyType).Interface()
			c.BodyParser(body)
			response := callHandlerWithBody(handler, body)
			return c.JSON(response)
		}
		response := callHandler(handler)
		return c.JSON(response)
	}
}

// Determines the request body struct type
func getHandlerBodyType(handler any) reflect.Type {
	handlerType := reflect.TypeOf(handler)
	if handlerType.NumIn() == 1 {
		return handlerType.In(0)
	}
	return nil
}

// Calls a handler function dynamically (no body)
func callHandler(handler any) any {
	handlerValue := reflect.ValueOf(handler)
	return handlerValue.Call([]reflect.Value{})[0].Interface()
}

// Calls a handler function dynamically (with body)
func callHandlerWithBody(handler any, body any) any {
	handlerValue := reflect.ValueOf(handler)
	return handlerValue.Call([]reflect.Value{reflect.ValueOf(body).Elem()})[0].Interface()
}
