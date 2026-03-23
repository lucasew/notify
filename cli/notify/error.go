package main

func reportError(err error) {
	log := New("error")
	log.Error(err)
}
