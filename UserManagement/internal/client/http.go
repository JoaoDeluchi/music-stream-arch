package httpClient

type (
	HttpClient interface {
		Get()
		Create()
		
	}

	httpClient struct {

	}
)

func NewClientHttp() {

}