// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package product

import (
	"context"
	"github.com/regalius/clean-arch/sample-app/pkg/model/product"
	"sync"
)

var (
	lockRepositoryMockGetByFilter   sync.RWMutex
	lockRepositoryMockGetByID       sync.RWMutex
	lockRepositoryMockGetByIDRange  sync.RWMutex
	lockRepositoryMockInsertProduct sync.RWMutex
	lockRepositoryMockUpdateProduct sync.RWMutex
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked Repository
//         mockedRepository := &RepositoryMock{
//             GetByFilterFunc: func(ctx context.Context, filter Filter) ([]product.Product, error) {
// 	               panic("mock out the GetByFilter method")
//             },
//             GetByIDFunc: func(ctx context.Context, id int64) (product.Product, error) {
// 	               panic("mock out the GetByID method")
//             },
//             GetByIDRangeFunc: func(ctx context.Context, minID int64, maxID int64) ([]product.Product, error) {
// 	               panic("mock out the GetByIDRange method")
//             },
//             InsertProductFunc: func(ctx context.Context, product product.Product) (int, error) {
// 	               panic("mock out the InsertProduct method")
//             },
//             UpdateProductFunc: func(ctx context.Context, newProduct product.Product) (int, error) {
// 	               panic("mock out the UpdateProduct method")
//             },
//         }
//
//         // use mockedRepository in code that requires Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// GetByFilterFunc mocks the GetByFilter method.
	GetByFilterFunc func(ctx context.Context, filter Filter) ([]product.Product, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id int64) (product.Product, error)

	// GetByIDRangeFunc mocks the GetByIDRange method.
	GetByIDRangeFunc func(ctx context.Context, minID int64, maxID int64) ([]product.Product, error)

	// InsertProductFunc mocks the InsertProduct method.
	InsertProductFunc func(ctx context.Context, product product.Product) (int, error)

	// UpdateProductFunc mocks the UpdateProduct method.
	UpdateProductFunc func(ctx context.Context, newProduct product.Product) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetByFilter holds details about calls to the GetByFilter method.
		GetByFilter []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// GetByIDRange holds details about calls to the GetByIDRange method.
		GetByIDRange []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MinID is the minID argument value.
			MinID int64
			// MaxID is the maxID argument value.
			MaxID int64
		}
		// InsertProduct holds details about calls to the InsertProduct method.
		InsertProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Product is the product argument value.
			Product product.Product
		}
		// UpdateProduct holds details about calls to the UpdateProduct method.
		UpdateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NewProduct is the newProduct argument value.
			NewProduct product.Product
		}
	}
}

// GetByFilter calls GetByFilterFunc.
func (mock *RepositoryMock) GetByFilter(ctx context.Context, filter Filter) ([]product.Product, error) {
	if mock.GetByFilterFunc == nil {
		panic("RepositoryMock.GetByFilterFunc: method is nil but Repository.GetByFilter was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	lockRepositoryMockGetByFilter.Lock()
	mock.calls.GetByFilter = append(mock.calls.GetByFilter, callInfo)
	lockRepositoryMockGetByFilter.Unlock()
	return mock.GetByFilterFunc(ctx, filter)
}

// GetByFilterCalls gets all the calls that were made to GetByFilter.
// Check the length with:
//     len(mockedRepository.GetByFilterCalls())
func (mock *RepositoryMock) GetByFilterCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	lockRepositoryMockGetByFilter.RLock()
	calls = mock.calls.GetByFilter
	lockRepositoryMockGetByFilter.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, id int64) (product.Product, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	lockRepositoryMockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockGetByID.RLock()
	calls = mock.calls.GetByID
	lockRepositoryMockGetByID.RUnlock()
	return calls
}

// GetByIDRange calls GetByIDRangeFunc.
func (mock *RepositoryMock) GetByIDRange(ctx context.Context, minID int64, maxID int64) ([]product.Product, error) {
	if mock.GetByIDRangeFunc == nil {
		panic("RepositoryMock.GetByIDRangeFunc: method is nil but Repository.GetByIDRange was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		MinID int64
		MaxID int64
	}{
		Ctx:   ctx,
		MinID: minID,
		MaxID: maxID,
	}
	lockRepositoryMockGetByIDRange.Lock()
	mock.calls.GetByIDRange = append(mock.calls.GetByIDRange, callInfo)
	lockRepositoryMockGetByIDRange.Unlock()
	return mock.GetByIDRangeFunc(ctx, minID, maxID)
}

// GetByIDRangeCalls gets all the calls that were made to GetByIDRange.
// Check the length with:
//     len(mockedRepository.GetByIDRangeCalls())
func (mock *RepositoryMock) GetByIDRangeCalls() []struct {
	Ctx   context.Context
	MinID int64
	MaxID int64
} {
	var calls []struct {
		Ctx   context.Context
		MinID int64
		MaxID int64
	}
	lockRepositoryMockGetByIDRange.RLock()
	calls = mock.calls.GetByIDRange
	lockRepositoryMockGetByIDRange.RUnlock()
	return calls
}

// InsertProduct calls InsertProductFunc.
func (mock *RepositoryMock) InsertProduct(ctx context.Context, product product.Product) (int, error) {
	if mock.InsertProductFunc == nil {
		panic("RepositoryMock.InsertProductFunc: method is nil but Repository.InsertProduct was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Product product.Product
	}{
		Ctx:     ctx,
		Product: product,
	}
	lockRepositoryMockInsertProduct.Lock()
	mock.calls.InsertProduct = append(mock.calls.InsertProduct, callInfo)
	lockRepositoryMockInsertProduct.Unlock()
	return mock.InsertProductFunc(ctx, product)
}

// InsertProductCalls gets all the calls that were made to InsertProduct.
// Check the length with:
//     len(mockedRepository.InsertProductCalls())
func (mock *RepositoryMock) InsertProductCalls() []struct {
	Ctx     context.Context
	Product product.Product
} {
	var calls []struct {
		Ctx     context.Context
		Product product.Product
	}
	lockRepositoryMockInsertProduct.RLock()
	calls = mock.calls.InsertProduct
	lockRepositoryMockInsertProduct.RUnlock()
	return calls
}

// UpdateProduct calls UpdateProductFunc.
func (mock *RepositoryMock) UpdateProduct(ctx context.Context, newProduct product.Product) (int, error) {
	if mock.UpdateProductFunc == nil {
		panic("RepositoryMock.UpdateProductFunc: method is nil but Repository.UpdateProduct was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		NewProduct product.Product
	}{
		Ctx:        ctx,
		NewProduct: newProduct,
	}
	lockRepositoryMockUpdateProduct.Lock()
	mock.calls.UpdateProduct = append(mock.calls.UpdateProduct, callInfo)
	lockRepositoryMockUpdateProduct.Unlock()
	return mock.UpdateProductFunc(ctx, newProduct)
}

// UpdateProductCalls gets all the calls that were made to UpdateProduct.
// Check the length with:
//     len(mockedRepository.UpdateProductCalls())
func (mock *RepositoryMock) UpdateProductCalls() []struct {
	Ctx        context.Context
	NewProduct product.Product
} {
	var calls []struct {
		Ctx        context.Context
		NewProduct product.Product
	}
	lockRepositoryMockUpdateProduct.RLock()
	calls = mock.calls.UpdateProduct
	lockRepositoryMockUpdateProduct.RUnlock()
	return calls
}