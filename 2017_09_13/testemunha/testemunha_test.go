package testemunha_test

import (
	. "github.com/globocom/dojo/2017_09_13/testemunha"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testemunha", func() {
	Context("Alan T. Uring na Globo.com com o trebuchet (4/11/4)", func() {
		crime := Crime{4, 11, 4}
		It("retorna 0 quando a teoria está certa", func() {
			Expect(crime.Testemunha(4, 11, 4)).To(Equal(0))
		})
		It("retorna 1 quando o assassino está incorreto", func() {
			Expect(crime.Testemunha(6, 11, 4)).To(Equal(1))
		})
		It("retorna 2 quando o local está incorreto", func() {
			Expect(crime.Testemunha(4, 5, 4)).To(Equal(2))
		})
		It("retorna 3 quando a arma está incorreta", func() {
			Expect(crime.Testemunha(4, 11, 6)).To(Equal(3))
		})
		It("não retorna 0 ou 1 quando a arma e local está incorreta", func() {
			hipotese := crime.Testemunha(4, 10, 5)
			Expect(hipotese).To(Not(Equal(1)))
			Expect(hipotese).To(Not(Equal(0)))
		})
		It("não retorna 0 nem 3 quando suspeito e local estão incorretos", func() {
			hipotese := crime.Testemunha(3, 10, 4)
			Expect(hipotese).To(Not(Equal(0)))
			Expect(hipotese).To(Not(Equal(3)))
		})
		It("não retorna 0 nem 2 quando suspeito e arma estão incorretos", func() {
			hipotese := crime.Testemunha(3, 11, 3)
			Expect(hipotese).To(Not(Equal(0)))
			Expect(hipotese).To(Not(Equal(2)))
		})
		It("não retorna 0 quando toda a teoria está incorreta", func() {
			Expect(crime.Testemunha(5, 10, 3)).To(Not(Equal(0)))
		})
	})
})

// Crime: 4/11/4
// Suspeito: 6
// Local: 11
// Arma: 6
var _ = Describe("Detetive", func() {
	It("Detetive deve resolver arma", func() {
		crime := Crime{4, 11, 4}
		Expect(Detetive(crime).Arma).To(Equal(4))
	})
	It("Detetive deve resolver local", func() {
		crime := Crime{4, 11, 4}
		Expect(Detetive(crime).Local).To(Equal(11))
	})
	It("Detetive deve resolver suspeito", func() {
		crime := Crime{4, 11, 4}
		Expect(Detetive(crime).Suspeito).To(Equal(4))
	})
})
