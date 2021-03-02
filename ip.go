package kata

import (
        "fmt"
        "strconv"
        "strings"
)

func Is_valid_ip(ip string) bool {
        check := strings.Split(ip, ".")
        if len(check) != 4 {
                return false
        }
        for i := 0; i < len(check); i++ {
                comparison, err := strconv.Atoi(check[i])
                fmt.Println(comparison)
                if err != nil {
                        return false
                }
                if comparison >= 0 && comparison <= 255 {
                        fmt.Println("valid")
                } else {
                        return false
                }
        }
        return true
}
