package router

// func TestGetAccount(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	manager := mock.NewMockManager(controller)
// 	r := mux.NewRouter()
// 	n := negroni.New()
// 	MakeBookHandlers(r, *n, manager)
// 	path, err := r.GetRoute("getBook").GetPathTemplate()
// 	assert.Nil(t, err)
// 	assert.Equal(t, "/v1/book/{id}", path)
// 	b := book.NewFixtureBook()
// 	manager.EXPECT().
// 			Get(b.ID).
// 			Return(b, nil)
// 	handler := getBook(manager)
// 	r.Handle("/v1/book/{id}", handler)
// 	ts := httptest.NewServer(r)
// 	defer ts.Close()
// 	res, err := http.Get(ts.URL + "/v1/book/" + b.ID.String())
// 	assert.Nil(t, err)
// 	assert.Equal(t, http.StatusOK, res.StatusCode)
// 	var d *book.Book
// 	json.NewDecoder(res.Body).Decode(&d)
// 	assert.NotNil(t, d)
// 	assert.Equal(t, b.ID, d.ID)
// }
