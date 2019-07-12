import QtQuick 2.10				//Item

Rectangle {
    id: mainWindow
    color: "#080806"
    width: 1024; height: 600

    Text {
        text: "Hello from QML"
        font.pixelSize: 24
        color: "#ffffff"
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
    }
}