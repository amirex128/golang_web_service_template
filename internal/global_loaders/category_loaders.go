package global_loaders

import (
	"github.com/amirex128/selloora_backend/internal/providers/pool"
)

var (
	categoryPool pool.Interface
)

//func initCategoryLoader() {
//	categoryPool = pool.NewPool(categoryTypeMapPool(), memory.NewMemoryPool(), 24*time.Hour, 20*time.Second, 4)
//	categoryPool.Start(context.Background())
//	<-categoryPool.Notify()
//	logrus.Info("category pool initialized")
//}
//
//func categoryTypeMapPool() func(ctx context.Context) (map[string]kv.Serializable, error) {
//	return func(ctx context.Context) (map[string]kv.Serializable, error) {
//		categories, err := models.NewMainManager().GetLevel1Categories()
//		b := make(map[string]kv.Serializable)
//		if err != nil {
//			return b, err
//		}
//		for i := range categories {
//			b[fmt.Sprintf("CACHE_%s_%d", "CATEGORY", categories[i].ID)] = categories[i]
//		}
//		return b, nil
//	}
//}
