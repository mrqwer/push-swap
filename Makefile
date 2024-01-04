.PHONY: push_swap checker clean

COMMON_FILES = main.go general.go init.go push.go reverse_rotate.go rotate.go swap.go utils.go


push_swap:
	@go build -tags "push_swap" -o push-swap $(COMMON_FILES) push_swap.go sort_complex.go sort_small.go

checker:
	@go build -tags "checker" -o checker $(COMMON_FILES) checker.go

clean:
	@rm -f push-swap checker
