<?php 

// In your "php.ini" file, search for the file_uploads directive, and set it to On:
?>
<!DOCTYPE html>
<html>
<body>

<form action="_2_file_upload_script.php" method="post" enctype="multipart/form-data">
  Select image to upload:
  <input type="file" name="fileToUpload" id="fileToUpload">
  <input type="submit" value="Upload Image" name="submit">
</form>

</body>
</html>
