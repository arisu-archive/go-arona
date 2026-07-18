# Changelog

## [1.4.0](https://github.com/arisu-archive/go-arona/compare/v1.3.0...v1.4.0) (2026-07-18)


### Features

* add summon character  functionality ([52a9f66](https://github.com/arisu-archive/go-arona/commit/52a9f66ce4da83aed29e9f4b3dca7e16fd1eda04))
* **arena:** add daily and cumulative time reward functionalities ([34a16dc](https://github.com/arisu-archive/go-arona/commit/34a16dc1327830405f9e03ffdbec70ff9cf7c654))
* **friend:** add bulk retrieval function ([b29fe26](https://github.com/arisu-archive/go-arona/commit/b29fe267c2b908f2f2ca31dc7a3513ba86e4e070))
* **shop:** implement merchandise purchasing functionality ([2f559a7](https://github.com/arisu-archive/go-arona/commit/2f559a7794f1f6a6d17bbdf3518f70451bdc977c))

## [1.3.0](https://github.com/arisu-archive/go-arona/compare/v1.2.0...v1.3.0) (2026-07-17)


### Features

* add arena service ([9119732](https://github.com/arisu-archive/go-arona/commit/911973298ddaf1654e34abb26232d0df25c7c5fb))
* add custom header support ([9acadd8](https://github.com/arisu-archive/go-arona/commit/9acadd80552dcd6e2fd5247e88dcd7c8fd6ed200))
* add gateway bypass option to API requests and update related methods ([77c7e68](https://github.com/arisu-archive/go-arona/commit/77c7e68158a86acdae81cf5d696fd731d0bcd164))
* add initial implementation of Arona package with encoder and processor ([f3fa304](https://github.com/arisu-archive/go-arona/commit/f3fa3046f1bc219d82d1be76ff8ea2e0e88a4566))
* add service token for authentication in encoder requests ([4b29ffd](https://github.com/arisu-archive/go-arona/commit/4b29ffdf2e32f1f945090a3b5d3a784fb48f92af))
* add WithBossGroup method to request builder ([c18ca16](https://github.com/arisu-archive/go-arona/commit/c18ca166c196305014960b6dd89ef252ff3ba64d))
* allow runtime switching server ([1c8d5bd](https://github.com/arisu-archive/go-arona/commit/1c8d5bd90e4b4171febae427f9324de622f08d64))
* apply encryption handling ([68d809a](https://github.com/arisu-archive/go-arona/commit/68d809ab30b17bfac564de303b2a828a132b8a21))
* implement cafe service ([fad17f5](https://github.com/arisu-archive/go-arona/commit/fad17f569b715da4e2ab78069f76eca6652dbf3d))
* Implement encoder and services ([8e109da](https://github.com/arisu-archive/go-arona/commit/8e109da74379c2de903a5a4ec44e4cf443666539))


### Bug Fixes

* add nil checks for session keys in encryption ([32c04c6](https://github.com/arisu-archive/go-arona/commit/32c04c64f3aee9b34d31efef555f356221a5fd09))
* cannot set asia url ([1cdf2d7](https://github.com/arisu-archive/go-arona/commit/1cdf2d7bf07224d3d9b84df6bf5039d9d6e6f0e0))
* change ErrInvalidSession methods to value receivers ([09a3e93](https://github.com/arisu-archive/go-arona/commit/09a3e9371030bb38fd3f0c1bd62944f7390b5ad5))
* change publicKey type to pointer in Client and rsaEncrypt function ([7aef9ea](https://github.com/arisu-archive/go-arona/commit/7aef9ea90174886006f3a4bf1b24c68b9f245db4))
* **deps:** update module github.com/arisu-archive/arona-flatbuffers to v0.4.0 ([#9](https://github.com/arisu-archive/go-arona/issues/9)) ([05fe0d9](https://github.com/arisu-archive/go-arona/commit/05fe0d91a113ea196ba97e277b1b6e043ee791e1))
* **deps:** update module github.com/arisu-archive/arona-flatbuffers to v0.5.0 ([#13](https://github.com/arisu-archive/go-arona/issues/13)) ([fb4f32a](https://github.com/arisu-archive/go-arona/commit/fb4f32a2650b88cb07c842a74065995bdd02d315))
* **deps:** update module github.com/arisu-archive/arona-flatbuffers to v0.6.0 ([5ba0b53](https://github.com/arisu-archive/go-arona/commit/5ba0b53e2b811c125d3aec50ef3abdd112d5edd4))
* **deps:** update module github.com/arisu-archive/arona-protos to v1.0.2 ([#12](https://github.com/arisu-archive/go-arona/issues/12)) ([16a8436](https://github.com/arisu-archive/go-arona/commit/16a8436206d9c9c821bc8b3d44a1995131d3e11f))
* **deps:** update module github.com/arisu-archive/arona-protos to v1.4.1 ([#24](https://github.com/arisu-archive/go-arona/issues/24)) ([1ba594c](https://github.com/arisu-archive/go-arona/commit/1ba594c87a49644e89b64e3bcc057a1b3a516e2a))
* **deps:** update module github.com/arisu-archive/arona-protos to v1.4.3 ([#25](https://github.com/arisu-archive/go-arona/issues/25)) ([23f0509](https://github.com/arisu-archive/go-arona/commit/23f0509f9b4f2b82b366bd79d087167f6bfa7a78))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.27.1 ([#1](https://github.com/arisu-archive/go-arona/issues/1)) ([b04b051](https://github.com/arisu-archive/go-arona/commit/b04b051e9a76ee9c86f3698e71dfa8e1428f7a89))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.27.3 ([#7](https://github.com/arisu-archive/go-arona/issues/7)) ([39c0655](https://github.com/arisu-archive/go-arona/commit/39c0655d58dc6e1cf56e915650552995e718a8ec))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.28.3 ([#22](https://github.com/arisu-archive/go-arona/issues/22)) ([4dcbe04](https://github.com/arisu-archive/go-arona/commit/4dcbe043286114558af5d2514d3808a5cd03db12))
* **deps:** update module github.com/onsi/gomega to v1.38.3 ([#8](https://github.com/arisu-archive/go-arona/issues/8)) ([1bf3f0b](https://github.com/arisu-archive/go-arona/commit/1bf3f0bc8197c0d25230c1e4b56dc3718ff200fe))
* handle empty session key in WithSessionKey function ([25be9e1](https://github.com/arisu-archive/go-arona/commit/25be9e137d2287752726a4bdaa4d7f1e7567e19a))
* incorrect test data ([1950cb8](https://github.com/arisu-archive/go-arona/commit/1950cb89bb5737ce60fe78b1837eca840cb957cf))
* invalid payload format ([c25e23a](https://github.com/arisu-archive/go-arona/commit/c25e23a562e03563709f39629a153c7a7d6517d5))
* missing decode base64 in encrypted response ([3f92e05](https://github.com/arisu-archive/go-arona/commit/3f92e054a22ca147e07604d4b7d67c433511ec58))
* missing hash in guest packet ([badceaa](https://github.com/arisu-archive/go-arona/commit/badceaa82368b1c46b32d160c9e25843be182209))
* missing packet unmarhsalling logic ([ae04651](https://github.com/arisu-archive/go-arona/commit/ae04651441f37ed5316e6c54a85deb4cdbbd3057))
* missing request header in encoding request ([57b5680](https://github.com/arisu-archive/go-arona/commit/57b5680053140ce360de6543755adfb2da1050ed))
* missing server and encoder token in copy method ([7936cc0](https://github.com/arisu-archive/go-arona/commit/7936cc0d46cd04f807521c2dcd9734254bcd01ff))
* rename clan member list request for consistency ([d3eea7a](https://github.com/arisu-archive/go-arona/commit/d3eea7a01f28065f14f7802ba8dc101c9ec85e80))
* simplify session key check in withSessionKey function ([9f40d49](https://github.com/arisu-archive/go-arona/commit/9f40d4961ecf1372a482c5e044f04948ebf041ae))
* update error handling methods for consistency ([9a76aaa](https://github.com/arisu-archive/go-arona/commit/9a76aaa051097d01c08a5474f72eb58258744367))
* update error protocol handling in API response ([af2fbe5](https://github.com/arisu-archive/go-arona/commit/af2fbe51a7046bb36bed59bd248b06a8dfa461d7))
* update UserSession key bundles to use value types instead of pointers ([aac938f](https://github.com/arisu-archive/go-arona/commit/aac938fe0432338d998b65cfb9902b877583bf20))


### Code Refactoring

* align jp version encryption method ([c61972a](https://github.com/arisu-archive/go-arona/commit/c61972a78924d80ec8c01809fe2fdc3782f76bf0))
* **client:** improve client copy method to share http.Client configuration ([784e841](https://github.com/arisu-archive/go-arona/commit/784e84146dffaef191744560704be91af5298a94))
* make the user credential mutable ([71929c4](https://github.com/arisu-archive/go-arona/commit/71929c4f704a2f8fe1156340c3a48e8df818c361))
* update protocol encoder configuration to use a struct ([9114539](https://github.com/arisu-archive/go-arona/commit/9114539ed43267f440715b59fbad97e5af5e19a5))


### Miscellaneous Chores

* allow override request default setting ([99d6c30](https://github.com/arisu-archive/go-arona/commit/99d6c309201b21b363d673b165baad0604c4eeb7))
* **ci:** setup release please pipeline ([1fd3fac](https://github.com/arisu-archive/go-arona/commit/1fd3facac6960788ccf2ed4d6e88ac960b93fbfa))
* **dep:** bump arona-protos version ([02e4ba7](https://github.com/arisu-archive/go-arona/commit/02e4ba72da7c0e996784d2d5aa71d70e352a4ade))
* **dep:** bump arona-protos version ([e0cbbf9](https://github.com/arisu-archive/go-arona/commit/e0cbbf9630ff0de72cf738cc17e73cb35b39a224))
* **dep:** bump go version ([edbcabc](https://github.com/arisu-archive/go-arona/commit/edbcabc89638e3f8d255a3b34244dc90e3e986c1))
* **dep:** bump plana-protos version to v1.0.0 ([4e89484](https://github.com/arisu-archive/go-arona/commit/4e894841ffc158e08ce3d250516532babbfe95f9))
* **dep:** maintain go mod ([e85d584](https://github.com/arisu-archive/go-arona/commit/e85d5843ba0ea58e61af0124d314f56cd5733b2c))
* **dep:** maintain lock file ([ae0a381](https://github.com/arisu-archive/go-arona/commit/ae0a3814e8e1d5764278e55b54dd0649ad089b95))
* **deps:** update dependency go to v1.25.4 ([#5](https://github.com/arisu-archive/go-arona/issues/5)) ([8588acb](https://github.com/arisu-archive/go-arona/commit/8588acb3509b64d0467976cb3af9cdd9528a77bd))
* **deps:** update dependency go to v1.25.5 ([#6](https://github.com/arisu-archive/go-arona/issues/6)) ([e078f2b](https://github.com/arisu-archive/go-arona/commit/e078f2baab8237ac72fe8ef421071eab47fb7d53))
* **deps:** update dependency go to v1.25.6 ([#14](https://github.com/arisu-archive/go-arona/issues/14)) ([67e2058](https://github.com/arisu-archive/go-arona/commit/67e20584edbaa80ee47f9dd930858183b2c8b510))
* **deps:** update dependency go to v1.26.1 ([3e66f32](https://github.com/arisu-archive/go-arona/commit/3e66f3215efb479180339bd378ca62dd1a3b11b1))
* **deps:** update dependency go to v1.26.2 ([#20](https://github.com/arisu-archive/go-arona/issues/20)) ([7b37659](https://github.com/arisu-archive/go-arona/commit/7b37659c6a5b9be0ffc7964f4138d0b2298a15d7))
