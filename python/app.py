import urllib.request
import ssl

# This restores the same behavior as before.
context = ssl._create_unverified_context()

with urllib.request.urlopen('https://www.maxiquim.com.br/', context=context) as response:
   html = response.read()

print(html)

