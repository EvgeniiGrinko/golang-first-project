PATH=users_pages

prerun:
	[ -d $(PATH) ] && echo "Directory "$(PATH)" exists." ||  /bin/mkdir $(PATH) && /bin/chmod -R 775 $(PATH)


