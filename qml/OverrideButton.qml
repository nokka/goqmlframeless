import QtQuick 2.12
import QtQuick.Controls 2.3

Button {
    property alias cursorShape: mouseArea.cursorShape

    MouseArea
    {
        id: mouseArea
        anchors.fill: parent
        onPressed:  mouse.accepted = false
    }
}