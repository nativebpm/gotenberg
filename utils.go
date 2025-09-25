package gotenberg

// These functions help with optional parameters in options structs

// Predefined paper sizes for convenience
var (
	// Standard paper sizes (width x height in inches)
	PaperSizeLetter  = [2]float64{8.5, 11}     // Letter - 8.5 x 11 (default)
	PaperSizeLegal   = [2]float64{8.5, 14}     // Legal - 8.5 x 14
	PaperSizeTabloid = [2]float64{11, 17}      // Tabloid - 11 x 17
	PaperSizeLedger  = [2]float64{17, 11}      // Ledger - 17 x 11
	PaperSizeA0      = [2]float64{33.1, 46.8}  // A0 - 33.1 x 46.8
	PaperSizeA1      = [2]float64{23.4, 33.1}  // A1 - 23.4 x 33.1
	PaperSizeA2      = [2]float64{16.54, 23.4} // A2 - 16.54 x 23.4
	PaperSizeA3      = [2]float64{11.7, 16.54} // A3 - 11.7 x 16.54
	PaperSizeA4      = [2]float64{8.27, 11.7}  // A4 - 8.27 x 11.7
	PaperSizeA5      = [2]float64{5.83, 8.27}  // A5 - 5.83 x 8.27
	PaperSizeA6      = [2]float64{4.13, 5.83}  // A6 - 4.13 x 5.83
)

// Helper functions for working with predefined paper sizes

// A4 returns A4 paper size option
func A4() ConvOption {
	return WithPaperSize(PaperSizeA4[0], PaperSizeA4[1])
}

// Letter returns Letter paper size option
func Letter() ConvOption {
	return WithPaperSize(PaperSizeLetter[0], PaperSizeLetter[1])
}
