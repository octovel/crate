.PHONY: mergeds, pushs, pushd

mergeds:
	git checkout stable
	git merge development

pushs:
	git push -u github stable
	git push -u codeberg stable

pushd:
	git push -u github development
	git push -u codeberg development
