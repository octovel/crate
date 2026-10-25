.PHONY: d, mergeds, pushs, pushd, quickbuild

d:
	git checkout development

mergeds:
	git checkout stable
	git merge development

pushs:
	git push -u github stable
	git push -u codeberg stable

pushd:
	git push -u github development
	git push -u codeberg development

quickbuild:
	@pwsh ./scripts/quickbuild.ps1
