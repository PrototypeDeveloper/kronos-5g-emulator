package command

import (
	"fmt"
	"strings"

	"kronos/models"
)

// init registers mapon, mapoff and mapget.
func init() {
	item := Arg{Name: "item", Kind: KindChoice, Optional: true, Choices: models.MapNames()}
	where := "; without an item: every item. On the CUI/GUI target() picks a subscriber, otherwise the node and every subscriber"
	Register(Spec{
		Name: "mapon", Args: []Arg{item}, Scope: ScopeNode | ScopeUE | ScopeByTarget, Where: InScenario | Interactive,
		Help: "turn the mapping (the auto-fix of sent signals) of an item on" + where,
		Run:  func(env Env, args []Value) error { return setMapping(env, args, true) },
	})
	Register(Spec{
		Name: "mapoff", Args: []Arg{item}, Scope: ScopeNode | ScopeUE | ScopeByTarget, Where: InScenario | Interactive,
		Help: "turn the mapping of an item off: its fields are sent as signal.txt has them" + where,
		Run:  func(env Env, args []Value) error { return setMapping(env, args, false) },
	})
	Register(Spec{
		Name: "mapget", Args: []Arg{item}, Scope: ScopeNode | ScopeUE | ScopeByTarget, Where: InScenario | Interactive,
		Help: "show whether the mapping of an item is on or off" + where,
		Run:  getMapping,
	})
}

// mapArg returns the items an optional item argument names (every item without one).
func mapArg(args []Value) (models.MapOff, string, error) {
	if len(args) == 0 {
		return models.MapAll, "every item", nil
	}
	it, ok := models.MapItemByName(args[0].Str)
	if !ok {
		return 0, "", fmt.Errorf("unknown mapping item %s (%s)", args[0].Str, strings.Join(models.MapNames(), ", "))
	}
	return it.Of(), it.String(), nil
}

// setMapping turns the items on or off and reports it.
func setMapping(env Env, args []Value, on bool) error {
	items, what, err := mapArg(args)
	if err != nil {
		return err
	}
	if err := env.SetMapping(items, on); err != nil {
		return err
	}
	env.Info(fmt.Sprintf("mapping %s: %s", onOff(on), what))
	return nil
}

// onOff returns "on" or "off".
func onOff(on bool) string {
	if on {
		return "on"
	}
	return "off"
}

// mapState returns on, off, or (on the console) mixed with how many have it off.
func mapState(env Env, it models.MapItem) string {
	off, of := env.Mapping(it)
	switch off {
	case 0:
		return "on"
	case of:
		return "off"
	}
	return fmt.Sprintf("mixed (off on %d of %d)", off, of)
}

// getMapping reports one item, or the items that are not on.
func getMapping(env Env, args []Value) error {
	if len(args) > 0 {
		it, ok := models.MapItemByName(args[0].Str)
		if !ok {
			_, _, err := mapArg(args)
			return err
		}
		env.Info(it.String() + ": " + mapState(env, it))
		return nil
	}
	var notOn []string
	for _, it := range models.MapItems() {
		if st := mapState(env, it); st != "on" {
			notOn = append(notOn, it.String()+" "+st)
		}
	}
	if len(notOn) == 0 {
		env.Info("mapping: every item on")
		return nil
	}
	env.Info("mapping: " + strings.Join(notOn, ", ") + "; the others on")
	return nil
}
