package acceptance

import (
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("disable_director_verifiers command", func() {
	var (
		server *ghttp.Server
	)

	BeforeEach(func() {
		server = createTLSServer()
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/v0/staged/director/verifiers/install_time"),
				ghttp.RespondWith(http.StatusOK, `{ "verifiers": [
							{ "type":"some-verifier-type", "enabled":true }
						]}`),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("PUT", "/api/v0/staged/director/verifiers/install_time/some-verifier-type"),
				ghttp.RespondWith(http.StatusOK, `{
							"type": "some-verifier-type",
							"enabled": false
						}`),
				ghttp.VerifyJSON(`{ "enabled": false }`),
			),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	It("disables any verifiers passed in if they exist", func() {
		command := exec.Command(pathToMain,
			"--target", server.URL(),
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"disable-director-verifiers",
			"--type", "some-verifier-type",
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(string(session.Out.Contents())).To(Equal(`Disabling Director Verifiers...

The following verifiers were disabled:
- some-verifier-type
`))
	})

	It("errors if any verifiers passed in don't exist", func() {
		command := exec.Command(pathToMain,
			"--target", server.URL(),
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"disable-director-verifiers",
			"--type", "some-verifier-type",
			"-t", "another-verifier-type",
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		Eventually(session).Should(gexec.Exit(1))

		Expect(string(session.Out.Contents())).To(Equal(`The following verifiers do not exist:
- another-verifier-type

No changes were made.

`))
	})
})
