package main

import (
	"reflex/utils"
	"fmt"
	"strings"
    "sort"
)

var (
    protos  map[string]string
    glosses map[string]string
)

/*  Parses the rule into:
    from:   the original sound sequence
    to:     the new sound reflex
    before: the environment preceding the original sound
    after:  the environment after the original sound */
func parse(rule string) (string, string, string, string) {
    arrow := strings.Split(rule, ">")
    from := arrow[0]
    slash := strings.Split(arrow[1], "/")
    to := slash[0]
    if len(slash) == 1 {
        return from, to, "", ""
    }
    enviro := strings.Split(slash[1], "_")
    before := enviro[0]
    after := enviro[1]
    return from, to, before, after
}

/* Checks if the preceding conditioning environment fits. 
    Option = True means preceding environment. False is following environment. */
func check(env string, rule string, affix func(string, string) bool) bool {
    if rule == "" {
        return true
    }
    if rule == "#" {
        return env == ""
    }
    var nat_class []string
    if !strings.Contains(rule, "{") {
        nat_class = []string{rule}
    } else {
        nat_class = strings.Split(strings.Trim(rule, "{}"), ",")
    }
    for _, sound := range nat_class {
        if affix(env, sound) {
            return true
        }
    }
    return false
}

/*  Applies the rule to the proto-forms. */
func apply(rule string) {
    from, to, before, after := parse(rule)
    fmt.Println(from, to,before,after)
    for greek, proto := range protos {
        if strings.Contains(proto, from) {
            flank := strings.SplitN(proto, from, 2)
            if check(flank[0], before, strings.HasSuffix) && check(flank[1], after, strings.HasPrefix) {
                protos[greek] = strings.Replace(proto, from, to, 1)
            }
        }
    }
}

func remainder() {
    success := []string{}
    order := make([]string, 0, len(protos))
    for proto := range protos {
        order = append(order, proto)
    }
    sort.Strings(order)
    for _, k := range order {
        v := protos[k]
        if k == v {
            success = append(success, k)
        } else {
            spaces := strings.Repeat(" ", 20 - len(v))
            fmt.Println("Proto:", "*" + v + spaces, "Greek:", k)
        }
    }
    fmt.Println(success)
}

func main() {
    protos = make(map[string]string)
    glosses = make(map[string]string)
    utils.Init_maps(protos, glosses, "../input.txt")
    remainder()
    rule := utils.Wait_user()
    rules := []string{}
    for rule != "quit" {
        fmt.Println(rule)
        if rule == "" {
            fmt.Println("Error listening to sentence.")
            return
        }
        rules = append(rules,rule)
        apply(rule)
        remainder()
        rule = utils.Wait_user()
    }
    fmt.Println(rules)
}
