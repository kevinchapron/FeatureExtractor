# Feature Extractor

[![Languages](https://img.shields.io/badge/languages-En-green.svg)]()
[![Licence Apache2](https://img.shields.io/hexpm/l/plug.svg)](http://www.apache.org/licenses/LICENSE-2.0)

---

This project is a **GoLang application** designed to be used on a RaspBerry Pi3.
It extracts features from datasets recorded with _Accelerometer/Gyroscope/Magnetometer_.
It is programmable to adapt itself to your own device (6DOF, 9DOF ...)

The program is able to extract 35 features per device linked. For example, A 9-DOF will have 105 features.
Features extracted are the ones below : 

|  Time Domain  | |Frequential Domain |
|---|---|---|
| Mean For each Axis & mean on all axis  |      | DC Component for each Axis |
| Standard Deviation for each Axis & on all axis  |      |  Energy for each Axis |
| Skewness for each Axis & on all axis  |      |  Entropy for each Axis |
| Kurtosis for each Axis & on all axis  |
| Corellation between : <ul><li>X-Axis and Y-Axis</li><li>X-Axis and Z-Axis</li><li>X-Axis and Total</li><li>Y-Axis and Z-Axis</li><li>Y-Axis and Total-Axis</li><li>Z-Axis and Total-Axis</li></ul>|
| Zero-Crossing-Rate for each Axis & on all axis  |
This program can save features extracted in a CSV-File aswell

More information
---

This project has been created to extract features of real-time data, in the **[LIARA](http://liara.uqac.ca/)** lab 
(Laboratoire d'Intelligence Ambiante pour la Reconnaissance d'Activités), at the 
« Université du Québec À Chicoutimi (**[UQAC](http://www.uqac.ca/)**) »

Author
---
**[Kévin CHAPRON](http://kevin-chapron.fr/)** - _2016_

Download
---
To download and use this project, just use the command below : 
> go get github.com/kevinchapron/FeatureExtractor

To download sources through git, just clone it using the command below : 
> git clone https://github.com/kevinchapron/FeatureExtractor.git

License
---
    Copyright 2016 Kévin Chapron

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
