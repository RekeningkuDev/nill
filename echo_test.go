package null

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type user struct {
	Name  Type[string] `json:"name" form:"name"`
	Email Type[string] `json:"email" form:"email"`
}

func TestFromQuery(t *testing.T) {

	t.Run("value not empty or null", func(t *testing.T) {

	})

	t.Run("one of the value is null", func(t *testing.T) {
	})

	t.Run("one of the value is empty", func(t *testing.T) {
	})

	t.Run("only has one value", func(t *testing.T) {
	})
}

func TestFromParam(t *testing.T) {

	t.Run("value not empty or null", func(t *testing.T) {

	})

	t.Run("one of the value is null", func(t *testing.T) {
	})

	t.Run("one of the value is empty", func(t *testing.T) {
	})

	t.Run("only has one value", func(t *testing.T) {
	})
}

func TestFromHeader(t *testing.T) {

	t.Run("value not empty or null", func(t *testing.T) {

	})

	t.Run("one of the value is null", func(t *testing.T) {
	})

	t.Run("one of the value is empty", func(t *testing.T) {
	})

	t.Run("only has one value", func(t *testing.T) {
	})
}

func TestFromXlm(t *testing.T) {

	t.Run("value not empty or null", func(t *testing.T) {

	})

	t.Run("one of the value is null", func(t *testing.T) {
	})

	t.Run("one of the value is empty", func(t *testing.T) {
	})

	t.Run("only has one value", func(t *testing.T) {
	})
}

func TestFromJson(t *testing.T) {
	t.Run("value not empty or null", func(t *testing.T) {
		userJSON := `{"name":"Jon Snow","email":"jon@labstack.com"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		u := new(user)
		if err := c.Bind(u); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, u.Name.Valid)
		assert.Equal(t, "Jon Snow", u.Name.Value)
		assert.True(t, u.Email.Valid)
		assert.Equal(t, "jon@labstack.com", u.Email.Value)
	})

	t.Run("one of the value is null", func(t *testing.T) {
		userJSON := `{"name":"Jon Snow","email":null}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		u := new(user)
		if err := c.Bind(u); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, u.Name.Valid)
		assert.Equal(t, "Jon Snow", u.Name.Value)
		assert.False(t, u.Email.Valid)
		assert.Equal(t, "", u.Email.Value)
	})

	t.Run("one of the value is empty", func(t *testing.T) {
		userJSON := `{"name":"Jon Snow","email":""}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		u := new(user)
		if err := c.Bind(u); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, u.Name.Valid)
		assert.Equal(t, "Jon Snow", u.Name.Value)
		assert.True(t, u.Email.Valid)
		assert.Equal(t, "", u.Email.Value)
	})

	t.Run("only has one value", func(t *testing.T) {
		userJSON := `{"name":"Jon Snow"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		u := new(user)
		if err := c.Bind(u); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, u.Name.Valid)
		assert.Equal(t, "Jon Snow", u.Name.Value)
		assert.False(t, u.Email.Valid)
		assert.Equal(t, "", u.Email.Value)
	})
}

func TestFromForm(t *testing.T) {

	t.Run("value not empty or null", func(t *testing.T) {

	})

	t.Run("one of the value is null", func(t *testing.T) {
	})

	t.Run("one of the value is empty", func(t *testing.T) {
	})

	t.Run("only has one value", func(t *testing.T) {
	})
}
