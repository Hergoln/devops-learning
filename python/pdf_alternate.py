import PyPDF2 as pypdf

original_filename = ""
output_filename = ""

with open(original_filename, 'rb') as original_pdf:
  reader = pypdf.PdfReader(original_pdf)

  first_page = reader.pages[0]

  writer = pypdf.PdfWriter()
  writer.add_page(first_page)

if original_filename is output_filename:
  print(f"original_filename ({original_filename}) is the same as output_filename, overwriting original file")

with open(output_filename, 'wb') as output_pdf:
  writer.write(output_pdf)
