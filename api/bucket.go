package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/smarthut/smarthut/model"
)

func (api *API) bucketCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bucketName := chi.URLParam(r, "bucket")
		bucket, err := model.GetBucket(api.db, bucketName)
		if err != nil {
			handleError(errors.Wrapf(err, "unable to find bucket with name %s", bucketName), w, r)
			return
		}
		ctx := context.WithValue(r.Context(), bucketKey, bucket)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) listBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := model.AllBuckets(api.db)
	if err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, buckets)
}

func (api *API) createBucket(w http.ResponseWriter, r *http.Request) {
	var data model.JSONBucket
	if err := render.DecodeJSON(r.Body, &data); err != nil {
		handleError(err, w, r)
		return
	}
	bucket, err := model.NewBucket(data.Name, data.Data)
	if err != nil {
		handleError(err, w, r)
		return
	}
	if err := api.db.Save(bucket); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, bucket)
}

func (api *API) getBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.Context().Value(bucketKey).(*model.JSONBucket)
	render.JSON(w, r, bucket.Data)
}

func (api *API) updateBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.Context().Value(bucketKey).(*model.JSONBucket)
	var data map[string]interface{}
	if err := render.DecodeJSON(r.Body, &data); err != nil {
		handleError(err, w, r)
		return
	}
	if err := bucket.UpdateData(api.db, data); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, bucket.Data)
}

func (api *API) deleteBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.Context().Value(bucketKey).(*model.JSONBucket)
	if err := bucket.Delete(api.db); err != nil {
		handleError(err, w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("bucket %s has been removed", bucket.Name)))
}
