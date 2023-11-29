package omalley

import (
	"strings"

	"github.com/aziflaj/caboose/megahash"
	"github.com/aziflaj/caboose/sarge"
)

func Execute(store *megahash.MegahashTable, command string, args []string) string {
	switch strings.ToUpper(command) {
	case "PING":
		return ping(args)
	case "ECHO":
		return echo(args)
	case "SET":
		return set(store, args)
	case "GET":
		return get(store, args)
	case "DELETE":
		return delete(store, args)
	case "EXISTS":
		return exists(store, args)
	case "KEYS":
		return keys(store, args)
	default:
		return unknownCommand(command)
	}
}

func ping(args []string) string {
	if len(args) == 0 {
		return sarge.SerializeBulkString("PONG")
	}
	return sarge.SerializeArray(args)
}

func echo(args []string) string {
	return sarge.SerializeArray(args)
}

func set(store *megahash.MegahashTable, args []string) string {
	if len(args) != 2 {
		return sarge.SerializeError("Wrong number of arguments for SET")
	}

	store.Set(args[0], args[1])

	return sarge.SerializeBulkString("OK")
}

func get(store *megahash.MegahashTable, args []string) string {
	if len(args) != 1 {
		return sarge.SerializeError("Wrong number of arguments for GET")
	}

	val := store.Get(args[0])

	return sarge.SerializeBulkString(val)
}

// TODO: you were too lazy to implement bools
func exists(store *megahash.MegahashTable, args []string) string {
	if len(args) != 1 {
		return sarge.SerializeError("Wrong number of arguments for EXISTS")
	}

	val := store.Exists(args[0])
	if val {
		return sarge.SerializeInteger(1)
	}

	return sarge.SerializeInteger(0)
}

func delete(store *megahash.MegahashTable, args []string) string {
	if len(args) != 1 {
		return sarge.SerializeError("Wrong number of arguments for DELETE")
	}

	store.Delete(args[0])

	return sarge.SerializeBulkString("OK")
}

func keys(store *megahash.MegahashTable, args []string) string {
	if len(args) != 0 {
		return sarge.SerializeError("Wrong number of arguments for KEYS")
	}

	keys := store.Keys()

	return sarge.SerializeArray(keys)
}

func unknownCommand(command string) string {
	return sarge.SerializeError("Unknown command: " + command)
}
