// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrderNewBulkParams creates a new OrderNewBulkParams object
// with the default values initialized.
func NewOrderNewBulkParams() *OrderNewBulkParams {
	var ()
	return &OrderNewBulkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderNewBulkParamsWithTimeout creates a new OrderNewBulkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderNewBulkParamsWithTimeout(timeout time.Duration) *OrderNewBulkParams {
	var ()
	return &OrderNewBulkParams{

		timeout: timeout,
	}
}

// NewOrderNewBulkParamsWithContext creates a new OrderNewBulkParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderNewBulkParamsWithContext(ctx context.Context) *OrderNewBulkParams {
	var ()
	return &OrderNewBulkParams{

		Context: ctx,
	}
}

// NewOrderNewBulkParamsWithHTTPClient creates a new OrderNewBulkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderNewBulkParamsWithHTTPClient(client *http.Client) *OrderNewBulkParams {
	var ()
	return &OrderNewBulkParams{
		HTTPClient: client,
	}
}

/*OrderNewBulkParams contains all the parameters to send to the API endpoint
for the order new bulk operation typically these are written to a http.Request
*/
type OrderNewBulkParams struct {

	/*Orders
	  An array of orders.

	*/
	Orders *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order new bulk params
func (o *OrderNewBulkParams) WithTimeout(timeout time.Duration) *OrderNewBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order new bulk params
func (o *OrderNewBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order new bulk params
func (o *OrderNewBulkParams) WithContext(ctx context.Context) *OrderNewBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order new bulk params
func (o *OrderNewBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order new bulk params
func (o *OrderNewBulkParams) WithHTTPClient(client *http.Client) *OrderNewBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order new bulk params
func (o *OrderNewBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrders adds the orders to the order new bulk params
func (o *OrderNewBulkParams) WithOrders(orders *string) *OrderNewBulkParams {
	o.SetOrders(orders)
	return o
}

// SetOrders adds the orders to the order new bulk params
func (o *OrderNewBulkParams) SetOrders(orders *string) {
	o.Orders = orders
}

// WriteToRequest writes these params to a swagger request
func (o *OrderNewBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Orders != nil {

		// form param orders
		var frOrders string
		if o.Orders != nil {
			frOrders = *o.Orders
		}
		fOrders := frOrders
		if fOrders != "" {
			if err := r.SetFormParam("orders", fOrders); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}