$texlive_packages = @(
    "multirow",
    "booktabs",
    "framed",
    "fvextra",
    "courier",
    "efbox",
    "grffile",
    "pdfpages",
    "tcolorbox",
    "wasysym",
    "wasy",
    "fancyvrb",
    "xcolor",
    "etoolbox",
    "upquote",
    "lineno",
    "eso-pic",
    "lstaddons",
    "pdflscape",
    "infwarerr",
    "pgf",
    "environ",
    "trimspaces",
    "listings",
    "pdfcol"
)

$extra_packages = $texlive_packages -join ','

$TEXLIVE_DIR = "D:/texlive/aio"

# choco install texlive --version=2022.20221202 --params "'/collections:pictures,latex,latexextra,latexrecommended'" --execution-timeout 5400
choco install texlive --version=2022.20221202 --params "'/collections:pictures,latex /InstallationPath:${TEXLIVE_DIR} /extraPackages:${extra_packages}'" -ErrorAction SilentlyContinue

$env:GENDOC_LATEXPATH = "${TEXLIVE_DIR}/bin/windows"
