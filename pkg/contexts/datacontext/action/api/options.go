package api

func NewOptions(olist ...Option) *Options {
	var opts Options
	opts.Priority = -1
	for _, o := range olist {
		o.ApplyActionHandlerOptionTo(&opts)
	}
	return &opts
}

func (o *Options) ApplyActionHandlerOptionTo(opts *Options) {
	if o.Action != "" {
		opts.Action = o.Action
	}
	if o.Priority > 0 {
		opts.Priority = o.Priority
	}
	opts.Selectors = append(opts.Selectors, o.Selectors...)
	opts.Versions = append(opts.Versions, o.Versions...)
}
