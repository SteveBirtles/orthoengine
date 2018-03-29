package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
)

func initiateAPI() {

	linkToLua(L, APILoadMap, "LoadMap")
	linkToLua(L, APIGetTile, "GetTile")
	linkToLua(L, APISetTile, "SetTile")
	linkToLua(L, APISetFocus, "SetFocus")
	linkToLua(L, APISetZoom, "SetZoom")

	linkToLua(L, APIGetId, "GetId")
	linkToLua(L, APIGetEntities, "GetEntities")
	linkToLua(L, APISetEntityPosition, "SetEntityPosition")
	linkToLua(L, APIGetEntityPosition, "GetEntityPosition")
	linkToLua(L, APISetEntityVelocity, "SetEntityVelocity")
	linkToLua(L, APIPathFind, "PathFind")
	linkToLua(L, APICreateEntity, "CreateEntity")
	linkToLua(L, APISetEntitySprite, "SetEntitySprite")
	linkToLua(L, APIGetEntitySprite, "GetEntitySprite")
	linkToLua(L, APISetEntityAnimation, "SetEntityAnimation")
	linkToLua(L, APISetEntityScript, "SetEntityScript")
	linkToLua(L, APIGetEntityScript, "GetEntityScript")
	linkToLua(L, APISetEntityProperty, "SetEntityProperty")
	linkToLua(L, APIGetEntityProperty, "GetEntityProperty")
	linkToLua(L, APIListEntityProperties, "ListEntityProperties")
	linkToLua(L, APIDeleteEntity, "DeleteEntity")
	linkToLua(L, APIEntityProximity, "EntityProximity")
	linkToLua(L, APISetFocusEntity, "SetFocusEntity")
	linkToLua(L, APIGetEntityClass, "GetEntityClass")
	linkToLua(L, APISetEntityClass, "SetEntityClass")
	linkToLua(L, APIOverrideClassScript, "OverrideClassScript")
	linkToLua(L, APISetEntityActive, "SetEntityActive")
	linkToLua(L, APISetClassActive, "SetClassActive")
	linkToLua(L, APISetAllActive, "SetAllActive")

	linkToLua(L, APIKeyPressed, "KeyPressed")
	linkToLua(L, APIKeyJustPressed, "KeyJustPressed")
	linkToLua(L, APIDisplayText, "DisplayText")
	linkToLua(L, APIDisplayOptions, "DisplayOptions")
	linkToLua(L, APIPlaySound, "PlaySound")
	linkToLua(L, APIPlayMusic, "PlayMusic")
	linkToLua(L, APIPauseMusic, "PauseMusic")
	linkToLua(L, APIEndGame, "EndGame")

	linkToLua(L, APIStartTimer, "StartTimer")
	linkToLua(L, APIGetTimer, "GetTimer")
	linkToLua(L, APICancelTimer, "CancelTimer")
	linkToLua(L, APISetPersistent, "SetPersistent")
	linkToLua(L, APIGetPersistent, "GetPersistent")

	linkToLua(L, luaPrint, "print")

}

//LoadMap(map)
func APILoadMap(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetTile(x, y, z) -> base, front
func APIGetTile(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetTile(x, y, z, base, front)
func APISetTile(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetFocus(x, y)
func APISetFocus(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetZoom(zoom)
func APISetZoom(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}

//GetId()
func APIGetId(L *lua.LState) int {
	L.Push(lua.LNumber(currentEntity))
	return 1
}
//GetEntities(x, y, z, radius) -> ids
func APIGetEntities(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityPosition(id, x, y, z)
func APISetEntityPosition(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetEntityPosition(id) -> x, y, z
func APIGetEntityPosition(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityVelocity(id, dx, dy, dz, n)
func APISetEntityVelocity(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//PathFind(id, targetid, searchdepth, velocity)
func APIPathFind(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//CreateEntity(x, y, z, sprite, class, script) -> id
func APICreateEntity(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntitySprite(id, sprite)
func APISetEntitySprite(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetEntitySprite(id) -> sprite
func APIGetEntitySprite(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityAnimation(id, first, last, speed)
func APISetEntityAnimation(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityScript(id, script)
func APISetEntityScript(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetEntityScript(id) -> script
func APIGetEntityScript(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityProperty(id, property, value)
func APISetEntityProperty(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetEntityProperty(id, property)
func APIGetEntityProperty(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//ListEntityProperties(id) -> map
func APIListEntityProperties(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//DeleteEntity(id)
func APIDeleteEntity(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//EntityProximity(id1, id2)
func APIEntityProximity(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetFocusEntity(id, follow)
func APISetFocusEntity(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetEntityClass(id)
func APIGetEntityClass(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityClass(id, class)
func APISetEntityClass(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//OverrideClassScript(class, script)
func APIOverrideClassScript(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetEntityActive(id, boolean, global)
func APISetEntityActive(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetClassActive(id, boolean, global)
func APISetClassActive(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//SetAllActive(boolean)
func APISetAllActive(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}

//KeyPressed(key) -> boolean
func APIKeyPressed(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//KeyJustPressed(key) -> boolean
func APIKeyJustPressed(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//DisplayText(text, image, justification)
func APIDisplayText(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//DisplayOptions(text, options) -> option
func APIDisplayOptions(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//PlaySound(sound)
func APIPlaySound(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//PlayMusic(music, looping)
func APIPlayMusic(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//PauseMusic()
func APIPauseMusic(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//EndGame(message)
func APIEndGame(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}

//StartTimer(name)
func APIStartTimer(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetTimer(name) -> milliseconds
func APIGetTimer(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//CancelTimer(name)
func APICancelTimer(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}

//SetPersistent(name, value)
func APISetPersistent(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
//GetPersistent(name) -> value
func APIGetPersistent(L *lua.LState) int {
	x := L.ToString(1)
	fmt.Println(x)
	return 0
}
