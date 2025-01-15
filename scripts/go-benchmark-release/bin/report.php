<?php

$versions = explode(" ", "v5.0.15 v5.1.2 v5.2.6 v5.3.9 master");

$status = [];
$resultKeys = array("name", "count", "duration", "_", "memory", "_", "allocs", "_");

foreach ($versions as $version) {
	$files = glob("out/".$version."/*.log");
	foreach ($files as $file) {
		$contents = array_map("trim", file($file));
		$pass = in_array("PASS", $contents);

		$valid = false;

		foreach ($contents as $line) {
			$match = strpos($line, "ns/op") !== false;
			if (!$match) {
				continue;
			}

			$valid = true;

			$res = preg_split("/[\s]+/", $line);
			$arr = array_combine($resultKeys, $res);

			$name = $arr['name'];
			unset($arr['_'], $arr['name']);
			$arr['pass'] = $pass;

			$status[$name][$version] = $arr;
		}

		if (!$valid) {
			echo "No data: ".$file."\n";
		}
	}
}

$header = ["Test name", ...$versions];
$keys = array_keys($status);
sort($keys);

function csv($in) {
	echo sprintf('=SPLIT("%s", ",")', implode(",", $in)) . "\n";
}

csv($header);

foreach ($keys as $key) {
	$row = array($key);
	foreach ($versions as $version) {
		$duration = 0;
		if (isset($status[$key][$version])) {
			$duration = $status[$key][$version]['duration'];
		}
		$row []= $duration;
	}

	csv($row);
}
