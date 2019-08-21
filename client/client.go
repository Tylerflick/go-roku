package client

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Tylerflick/go-roku/models"
)

const (
	Port = 8060
)

type Key struct {
	Name string
}

type Client struct {
	http    *http.Client
	address string
}

// NewClient ...
func NewClient(address string) *Client {
	c := Client{
		http:    &http.Client{},
		address: address,
	}
	return &c
}

// GetApps ...
func (c *Client) GetApps() (*models.Apps, error) {
	// query/apps
	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://%s:8060/query/apps", c.address))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apps models.Apps
	err = xml.Unmarshal(body, &apps)
	return &apps, err
}

// GetActiveApp ...
func (c *Client) GetActiveApp() (*models.ActiveApp, error) {
	// query/active-app
	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://%s:8060/query/active-app", c.address))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var app *models.ActiveApp
	err = xml.Unmarshal(body, app)
	return app, err
}

func (c *Client) GetInputs() []string {
	return []string{
		"InputTuner",
		"InputHDMI1",
		"InputHDMI2",
		"InputHDMI3",
		"InputHDMI4",
		"InputAV1",
	}
}

func (c *Client) VolumeDown() (int, error) {
	return c.KeyPress(&Key{Name: "VolumeDown"})
}

func (c *Client) VolumeMute() (int, error) {
	return c.KeyPress(&Key{Name: "VolumeMute"})
}

func (c *Client) VolumeUp() (int, error) {
	return c.KeyPress(&Key{Name: "VolumeUp"})
}

func (c *Client) PowerOff() (int, error) {
	return c.KeyPress(&Key{Name: "PowerOff"})
}

func (c *Client) ChannelUp() (int, error) {
	return c.KeyPress(&Key{Name: "ChannelUp"})
}

func (c *Client) ChannelDown() (int, error) {
	return c.KeyPress(&Key{Name: "ChannelDown"})
}

func (c *Client) Home() (int, error) {
	return c.KeyPress(&Key{Name: "home"})
}

// ChangeInput ...
func (c *Client) ChangeInput(source string) (int, error) {
	return c.KeyPress(&Key{Name: source})
}

// KeyDown ...
func (c *Client) KeyDown(key *Key) (int, error) {
	// keydown/key
	return c.post(fmt.Sprintf("http://%s:8060/keydown/key", c.address), "", nil)
}

// KeyUp ...
func (c *Client) KeyUp(key *Key) (int, error) {
	// keyup/key
	return c.post(fmt.Sprintf("http://%s:8060/keyup/%s", c.address, key.Name), "", nil)
}

// KeyPress ...
func (c *Client) KeyPress(key *Key) (int, error) {
	// keypress/key
	return c.post(fmt.Sprintf("http://%s:8060/keypress/%s", c.address, key.Name), "", nil)
}

// LaunchApp ...
func (c *Client) LaunchApp(id *string) {
	// TODO: Implement
}

func (c *Client) InstallApp(id *string) {
	// TODO: Implement
}

func (c *Client) GetDeviceInfo() {
	// query/device-info
	// TODO: Implement
}

func (c *Client) GetAppIcon() {
	// query/icon/appID
	// TODO: Implement
}

func (c *Client) post(url, form string, body io.Reader) (int, error) {
	resp, err := c.http.Post(url, form, body)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, err
}
