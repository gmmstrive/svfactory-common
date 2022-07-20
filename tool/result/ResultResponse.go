package result

type Response[T any] struct {
	Msg     string
	Code    int
	Data    T
	Success bool
}

func Ok(r IMessage) Response[any] {
	return Response[any]{Msg: r.GetMessage(), Code: r.GetCode(), Data: nil, Success: true}
}

// git config --global url."http://gitlab.cz.svfactory.com.cn:6180/".insteadof "https://gitlab.com/"

// gitlab.com/svf-go/svfactory-common-cloud/svfactory-common
