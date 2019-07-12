import QtQuick 2.10				//Item

Rectangle {
    id: mainWindow
    color: "#080806"
    width: 1024; height: 600

    OverrideButton{
        cursorShape: Qt.PointingHandCursor
    }

    MouseArea {
        width: 150
        height: 150
        hoverEnabled: true
        x: 50; y: 50
        cursorShape: Qt.PointingHandCursor
        Rectangle { anchors.fill: parent; color: "red"; }
        onPositionChanged: console.log("position", mouse.x, mouse.y)
        onContainsMouseChanged: console.log("containsMouse", containsMouse)
    }
}