
.PHONY: install
install: acc oj 

.PHONY: acc
acc: npm
	npm install -g atcoder-cli

.PHONY: oj
oj:
	pip3 install online-judge-tools

.PHONY: npm
npm: 
	curl -qL https://www.npmjs.com/install.sh | sh

.PHONY: login
login:
	acc login
	oj login https://atcoder.jp/

.PHONY: config
config:
	acc config default-task-choice all
	acc config default-template go


