# find text in currenty directory
grep "text to find" -R .

# find in *.pas text file and print line number
grep -an 'MasterManager' -R . --include='*.pas'
