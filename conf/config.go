package config

//Set the configuration values on this file

// Get gets you the value for the key
func Get(t string) string {
    switch t {
    case "WEBSERVER_PORT":
        return "9000"
    case "":
        return ""
    case "SESSION_NAME":
        return "master-session"
    case "REDIS_DB":
        return "0"
    case "REDIS_PASSWORD":
        return ""
    case "REDIS_HOST":
        return "localhost"
    case "REDIS_PORT":
        return "6379"
    default:
        return ""
    }
}
