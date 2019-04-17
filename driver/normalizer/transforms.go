package normalizer

import "github.com/bblfsh/sdk/v3/driver"

var Transforms = driver.Transforms{
	Namespace:      "ruby",
	Preprocess:     Preprocess,
	PreprocessCode: PreprocessCode,
	Normalize:      Normalize,
	Annotations:    Native,
}
