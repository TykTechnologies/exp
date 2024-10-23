<html>
<head>
  <title>{{ .title }}</title>
  <style type="text/css">
    a {
      display: inline-block;
      vertical-align: middle;
    }

    .card {
      background: #eee;
      margin-bottom: 10px;
      vertical-align: top;
    }

    .container {
      padding: 2rem 0rem;
    }

    .card-image {
      display: block;
      min-height: 240px;

      margin-top: 5px;
      margin-bottom: 10px;

      background: #eee;
      background-position: center center;
      background-repeat: no-repeat;
      background-size: contain;
    }
  </style>

  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KK94CHFLLe+nY2dmCWGMq91rCGa5gtU4mk92HdvYe+M/SXH301p5ILy+dN9+nJOZ" crossorigin="anonymous">

</head>
<body>

<div class="container">
<div class="row">

<div class="col-12">

<h1>{{ .title }}</h1>


<div class="row">
{{ range .images }}

<div class="col-3">
<div class="card w-100" style="display: inline-block">
  <div class="card-body">
    <a href="{{ .Basename }}" style="background-image: url({{.Basename}});" class="card-image"></a>
  </div>
</div>
</div>

{{ end }}
</div>


</div>

{{ if or .dirs .files }}

<h2>Files:</h2>

<div class="col-12">

<ul>
{{ range .dirs }}
<li><a href="{{ .Basename }}/">{{ .Basename }}/</a></li>
{{ end }}
{{ range .files }}
<li><a href="{{ .Basename }}">{{ .Basename }}</a></li>
{{ end }}
</ul>

</div>

{{ end }}

</div>
</div>

</body>
</html>