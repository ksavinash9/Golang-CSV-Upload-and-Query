package redis_test

import (
    "encoding/json"
    "reflect"
    "../models"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
    client := models.SetupRedis()

    BeforeEach(func() {
        client = models.SetupRedis()
        Expect(client.FlushDb().Err()).NotTo(HaveOccurred())
    })

    AfterEach(func() {
        Expect(client.Close()).NotTo(HaveOccurred())
    })

    Describe("CSV Database - Redis - ", func() {

        It("should Echo", func() {
            echo := client.Echo("hello")
            Expect(echo.Err()).To(HaveOccurred())
            Expect(echo.Val()).To(Equal("hello"))
        })

        It("should Ping", func() {
            ping := client.Ping()
            Expect(ping.Err()).NotTo(HaveOccurred())
            Expect(ping.Val()).To(Equal("PONG"))
        })

        It("should Select", func() {
            sel := client.Select(1)
            Expect(sel.Err()).NotTo(HaveOccurred())
            Expect(sel.Val()).To(Equal("OK"))
        })

        It("should be able to LPop", func() {
            rPush := client.RPush("list", "one")
            Expect(rPush.Err()).NotTo(HaveOccurred())
            rPush = client.RPush("list", "two")
            Expect(rPush.Err()).NotTo(HaveOccurred())
            rPush = client.RPush("list", "three")
            Expect(rPush.Err()).NotTo(HaveOccurred())

            lPop := client.LPop("list")
            Expect(lPop.Err()).NotTo(HaveOccurred())
            Expect(lPop.Val()).To(Equal("one"))

            lRange := client.LRange("list", 0, -1)
            Expect(lRange.Err()).NotTo(HaveOccurred())
            Expect(lRange.Val()).To(Equal([]string{"two", "three"}))
        })

        It("should be able LPush", func() {
            lPush := client.LPush("list", "World")
            Expect(lPush.Err()).NotTo(HaveOccurred())
            lPush = client.LPush("list", "Hello")
            Expect(lPush.Err()).NotTo(HaveOccurred())

            lRange := client.LRange("list", 0, -1)
            Expect(lRange.Err()).NotTo(HaveOccurred())
            Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))
        })

        It("should be able LRange", func() {
            lPush := client.LPush("list", "World")
            Expect(lPush.Err()).NotTo(HaveOccurred())
            lPush = client.LPush("list", "Hello")
            Expect(lPush.Err()).NotTo(HaveOccurred())

            lRange := client.LRange("list", 0, -1)
            Expect(lRange.Err()).NotTo(HaveOccurred())
            Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))
        })

        It("should be able to pipeline", func() {
            set := client.Set("key2", "hello2", 0)
            Expect(set.Err()).NotTo(HaveOccurred())
            Expect(set.Val()).To(Equal("OK"))

            pipeline := client.Pipeline()
            set = pipeline.Set("key1", "hello1", 0)
            get := pipeline.Get("key2")
            incr := pipeline.Incr("key3")
            getNil := pipeline.Get("key4")

            cmds, _ := pipeline.Exec()
            Expect(cmds).To(HaveLen(4))
            Expect(pipeline.Close()).NotTo(HaveOccurred())

            Expect(set.Err()).NotTo(HaveOccurred())
            Expect(set.Val()).To(Equal("OK"))

            Expect(get.Err()).NotTo(HaveOccurred())
            Expect(get.Val()).To(Equal("hello2"))

            Expect(incr.Err()).NotTo(HaveOccurred())
            Expect(incr.Val()).To(Equal(int64(1)))

            Expect(getNil.Val()).To(Equal(""))
        })
    })

    Describe("debugging", func() {

        It("should DebugObject", func() {
            debug := client.DebugObject("foo")
            Expect(debug.Err()).To(HaveOccurred())
            Expect(debug.Err().Error()).To(Equal("ERR no such key"))

            client.Set("foo", "bar", 0)
            debug = client.DebugObject("foo")
            Expect(debug.Err()).NotTo(HaveOccurred())
            Expect(debug.Val()).To(ContainSubstring(`serializedlength:4`))
        })

    })
})

type numberStruct struct {
    Number int
}

func (s *numberStruct) MarshalBinary() ([]byte, error) {
    return json.Marshal(s)
}

func (s *numberStruct) UnmarshalBinary(b []byte) error {
    return json.Unmarshal(b, s)
}

func deref(viface interface{}) interface{} {
    v := reflect.ValueOf(viface)
    for v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    return v.Interface()
}