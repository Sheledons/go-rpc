package connection

import (
	"go-rpc/com/github/sheledon/entity"
)

type Pipeline struct {
	inboundHandlers []RpcInboundHandler
	outboundHandlers []RpcOutboundHandler
}
func NewDefaultPipeline() *Pipeline {
	p := &Pipeline{
		inboundHandlers: make([]RpcInboundHandler,0),
		outboundHandlers: make([]RpcOutboundHandler,0),
	}
	p.initInboundHandlers()
	p.initOutboundHandlers()
	return p
}
func (p *Pipeline) initInboundHandlers() {
	p.addInboundHandler(NewDecodeHandler())
	p.addInboundHandler(NewDisPatchHandler())
}
func (p *Pipeline) initOutboundHandlers()  {
	p.addOutboundHandler(NewEncodeHandler())
}
func (p *Pipeline) addInboundHandler(handler RpcInboundHandler) {
	p.inboundHandlers = append(p.inboundHandlers,handler)
}
func (p *Pipeline) addOutboundHandler(handler RpcOutboundHandler){
	p.outboundHandlers = append(p.outboundHandlers,handler)
}
func (p *Pipeline) ProcessRead(ctx *ConnectContext) {
	for _,h := range p.inboundHandlers{
		h.Read(ctx)
	}
}
func (p *Pipeline) ProcessWrite(ctx *ConnectContext,msg *entity.RpcMessage) {
	for _,h := range p.outboundHandlers{
		h.Write(ctx,msg)
	}
}