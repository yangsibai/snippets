FROM: <http://apple.stackexchange.com/questions/151716/notification-settings-arent-being-retained>

    cd `getconf DARWIN_USER_DIR`
    rm -rf com.apple.notificationcenter 
    killall usernoted; killall NotificationCenter
