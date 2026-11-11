// Constants of package command.

package command

import "time"

// The command line.

// TargetKeyword prefixes a command line to select one subscriber:
// `target(IMSI).send(X)`. It cannot be used as a command name.
const TargetKeyword = "target"

// Parameter kinds.
const (
	KindSignal  Kind = iota // signal name; must exist in signal.txt
	KindState               // state name; must exist in scenario.txt
	KindFormat              // Printf format string literal
	KindValue               // literal, variable or builtin ($imsi ...)
	KindCommand             // command name (the argument of help)
	KindRate                // bit rate as iperf writes it: 100M, 1G, 500k (or bit/s as an int)
	KindChoice              // one of the words in Arg.Choices (ul / dl ...)
)

// Command scopes.
const (
	ScopeNode     Scope = 1 << iota // runs on the node (SCTP association)
	ScopeUE                         // runs on a subscriber; on the command line target() picks one, otherwise all
	ScopeConsole                    // acts on the pseudo-call as a whole (help, linkup); target() is not allowed
	ScopeBySignal                   // send: ScopeUE for a UE-associated signal, ScopeNode otherwise
)

// Places where a command may be written.
const (
	InScenario  Where = 1 << iota // inside handlers and entry blocks of scenario.txt
	Interactive                   // on the CUI/GUI command line
)

// Value types.
const (
	TypeNone Type = iota
	TypeInt
	TypeFloat
	TypeString
	TypeDuration
)

// Token kinds.
const (
	TokEOF      TokenKind = iota
	TokIdent              // name: send, NGSetupRequest, start, a
	TokBuiltin            // $imsi (Text holds "imsi")
	TokInt                // 1, -5, 208930000000001 (Text keeps leading zeros)
	TokFloat              // 1.1
	TokDuration           // 5s, 100ms, 1m30s
	TokString             // "test" (Text holds the unescaped value)
	TokLParen
	TokRParen
	TokLBrace
	TokRBrace
	TokComma
	TokDot
	TokColon
	TokAssign
	TokStar
	TokIllegal // Text holds the error message
)

// Argument expression kinds.
const (
	ExprIdent   ExprKind = iota // a name; the parameter kind decides whether it is a signal, state, variable or command
	ExprBuiltin                 // $name
	ExprLiteral                 // int, float, string or duration
)

// The snap format.

// allowedVerbs are the fmt verbs a snap format may use.
const allowedVerbs = "vTtbcdoOqxXUeEfFgGs"

// Limits of ping (the traffic/ping package checks the same at run time).
const (
	pingMaxCount    = 1_000_000
	pingMinInterval = time.Millisecond
	pingMaxInterval = time.Hour
	pingMaxSize     = 65000
)

// Limits of iperf and iperftcp (the others come from the traffic/iperf package).
const (
	iperfMaxPort = 65535 // the highest UDP/TCP port
)

// Limits of wait.
const (
	waitMaxTime = time.Hour // the longest wait
)

// Subscribers.
const (
	imsiDigits = 15 // the number of digits of an IMSI given to target()
)

// "Did you mean" hints: the edit distance accepted as a typo is one edit per
// suggestRunesPerEdit runes of the name, but at least suggestMinEdits.
const (
	suggestRunesPerEdit = 3 // runes of the name per accepted edit
	suggestMinEdits     = 2 // edits always accepted, even for a short name
)
