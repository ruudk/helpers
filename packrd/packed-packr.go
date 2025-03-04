// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "38207c3586161f8d509733b27cb4597c"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"219f70cdf30f6f12b6c726d0cb3c36b8": "1f8b08000000000000ff9490cb4ef3301085f7798aa9aafc6a16b137ff86aa170995dba61bd845883af124b67032953d2e1215ef8e720116b0e9caf2f89cef1ccf1ceed11dd187242966c575b44ec3232b8ee17961988f6129a5c69350efd1a3a8a895adf2afa5620cb28c75ad1cc91775b4fdcd3a2dc3e0950d4d8fc28cf86de9555799bd6a71ddaac0e8b30b0246b6eb87bcd558dbceb2a5ee41afff5ffdfb4d2e66c51deda8fa4968485325c837b2b16c6239047d97945f25c7f6229c9aec526f96e4799eec8971094fc606a8ad43b0015464ca1becd02b462d6047b027861b6d79b024ab740d357958980c6c07130fce906e92f91c8a432f30a2ff20a49bc35fc546c56d746e526503d5d6b03062475506678f1c7d07fd7235bd75d3fca30f59a5309c9f010000ffff3df549080c020000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("github.com/gobuffalo/helpers/genny/docs/templates", "../docs/templates")
		b.SetResolver("README.md.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "219f70cdf30f6f12b6c726d0cb3c36b8"})
		}()

	return nil
}()
