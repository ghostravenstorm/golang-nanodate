
package nanodate

import (
   "fmt"
   "log"
   "strconv"
   "errors"
)

var DebugLevel int = 0

/// Definition of what a Date contains
type Date struct {
   Milis   uint16 `json:"milis"`
   Seconds uint8  `json:"seconds"`
   Minutes uint8  `json:"minutes"`
   Hour    uint8  `json:"hour"`
   Day     uint8  `json:"day"`
   Month   uint8  `json:"month"`
   Year    uint16 `json:"year"`
}


/// Outputs all data fields to the console.
func (this *Date) Print() {

   fmt.Print("  Miliseconds:", this.Milis)
   fmt.Print("  Seconds:", this.Seconds)
   fmt.Print("  Minutes:", this.Minutes)
   fmt.Print("  Hour:", this.Hour)
   fmt.Print("  Day:", this.Day)
   fmt.Print("  Month:", this.Month)
   fmt.Println("  Year:", this.Year)
}

/// Returns several parts of a date formatted as strings.
/// Single digit numbers are preceded with a `0`.
func (this *Date) ToString() (string, string, string, string, string, string, string) {

   var yr string = strconv.Itoa(int(this.Year))
   var mo string = ""
   var da string = ""
   var hr string = ""
   var mt string = ""
   var sc string = ""
   var ms string = ""

   if this.Month < 10 {
      mo = "0" + strconv.Itoa(int(this.Month))
   } else {
      mo=strconv.Itoa(int(this.Month))
   }

   if this.Day < 10 {
      da = "0" + strconv.Itoa(int(this.Day))
   } else {
      da = strconv.Itoa(int(this.Day))
   }

   if this.Hour < 10 {
      hr = "0" + strconv.Itoa(int(this.Hour))
   } else {
      hr = strconv.Itoa(int(this.Hour))
   }

   if this.Minutes < 10 {
      mt = "0" + strconv.Itoa(int(this.Minutes))
   } else {
      mt = strconv.Itoa(int(this.Minutes))
   }

   if this.Seconds < 10 {
      sc = "0" + strconv.Itoa(int(this.Seconds))
   } else {
      sc = strconv.Itoa(int(this.Seconds))
   }

   if this.Milis < 100 {
      ms = "0" + strconv.Itoa(int(this.Milis))
      if this.Milis < 10 {
         ms = "00" + strconv.Itoa(int(this.Milis))
      }
   } else {
      ms = strconv.Itoa(int(this.Milis))
   }

   return yr, mo, da, hr, mt, sc, ms
}

/// Returns a string date based on the following formats
/// 0: yyyymmdd
/// 1: yyyy-mm-dd
/// 2: yyyymmddHHmm
/// 3: yyyy-mm-dd-HHmm
/// 4: yyyy
func (this *Date) ToStringWithFormat (formatType int) string {
   /// TODO:
   /// Needs to append a 0 if any of the months or days are single digit.
   /// Fix other cases with strconv.Atoi
   switch formatType {
      // case 0:{
      //    return string(d.Year) + string(d.Month) + string(d.Day)
      // }
      // case 1:{
      //    return strconv.Atoi(d.Year) + "-" +strconv.Atoi(d.Month) + "-" + strconv.Atoi(d.Day)
      // }
      // case 2:{
      //    return string(d.Year) + string(d.Month) + string(d.Day) + string(d.Hour) + string(d.Minutes)
      // }
      // case 3:{
      //    return string(d.Year) + "-" +string(d.Month) + "-" + string(d.Day) + "-" + string(d.Hour) + string(d.Minutes)
      // }
      // case 4:{
      //    return string(d.Year)
      // }
      default: {

         yr := strconv.Itoa((int(this.Year)))

         var mo string
         if this.Month >= 10 {
            mo = strconv.Itoa((int(this.Month)))
         } else {
            mo = "0" + strconv.Itoa((int(this.Month)))
         }

         var da string
         if this.Day >= 10 {
            da = strconv.Itoa((int(this.Day)))
         } else {
            da = "0" + strconv.Itoa((int(this.Day)))
         }

         return yr + "-" + mo + "-" + da
      }
   }
}

func (this *Date) ConvertFromDateStamp (dateStamp string) {
   /// WIP
}

/// Imports date from a string using the following formats:
///   "yyyymmddHHmm+"    Trunkcates any additional seconds.
///   "yyyymmddHHmm"
///   "yyyymmdd"
/// All other string forms will be rejected.
func (this *Date) ImportFromStringTypeA (str string) error {

   if DebugLevel >= 3 {
      log.Println("Parsing date", str)
   }

   /// Shared place holder for string to number converstion.
   var tempNumber int

   var err error

   if len(str) >= 12 {

      tempNumber, err = strconv.Atoi(str[0:4])
      this.Year = uint16(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[4:6])
      this.Month = uint8(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[6:8])
      this.Day = uint8(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[8:10])
      this.Hour = uint8(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[10:12])
      this.Minutes = uint8(tempNumber)
      if err != nil { return err }

   } else if len(str) == 8 {
      tempNumber, err = strconv.Atoi(str[0:4])
      this.Year = uint16(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[4:6])
      this.Month = uint8(tempNumber)
      if err != nil { return err }

      tempNumber, err = strconv.Atoi(str[6:8])
      this.Day = uint8(tempNumber)
      if err != nil { return err }
   } else {
      return errors.New("String to date formatting error.")
   }

   return nil
}

/// Returns true if this date is between two other given dates
func (this *Date) IsDateInRange (fromDate *Date, toDate *Date) bool {

   /// Check if any dates are valid.
   if fromDate == nil || toDate == nil {
      return true
   }

   /// Check if the date range should be ignored.
   if fromDate.IsValidAsEmpty() && toDate.IsValidAsEmpty() {
      return true
   }

   /// Return false if any field from any date contains a 0 month or a 0 day.
   if !this.IsValid() || !fromDate.IsValid() || !toDate.IsValid() {
      return false
   }

   var yr string
   var mo string
   var da string

   if DebugLevel == 1 {
      fmt.Printf("\n\n")
      log.Println("Checking date range.")
      log.Println("This date:")
      this.Print()
      fmt.Printf("\n")

      log.Println("From date:")
      fromDate.Print()
      fmt.Printf("\n")

      log.Println("To date:")
      toDate.Print()
      fmt.Printf("\n")
   }

   yr, mo, da, _, _, _, _ = this.ToString()
   nThisDate, _ := strconv.Atoi(yr + mo + da)

   yr, mo, da, _, _, _, _ = fromDate.ToString()
   nFromDate, _ := strconv.Atoi(yr + mo + da)

   yr, mo, da, _, _, _, _ = toDate.ToString()
   nToDate, _ := strconv.Atoi(yr + mo + da)

   if nThisDate >= nFromDate && nThisDate <= nToDate {
      if DebugLevel == 1 {
         log.Println("Date is in range.")
         fmt.Printf("\n")
      }
      return true
   } else {
      if DebugLevel == 1 {
         log.Println("Date is NOT in range.")
         fmt.Printf("\n")
      }
      return false
   }
}

func (this *Date) DoesDateMatch (date Date) bool {

   if date == Date{} {
      return true
   }

   if date.Year != 0 || date.Year != 0 {
      if this.Year != date.Year {
         return false
      }
   }

   if date.Month != 0 || date.Month != 0 {
      if this.Month != date.Month {
         return false
      }
   }

   if date.Day != 0 || date.Day != 0 {
      if this.Day != date.Day {
         return false
      }
   }

   if date.Hour != 0 || date.Hour != 0 {
      if this.Hour != date.Hour {
         return false
      }
   }

   if date.Minutes != 0 || date.Minutes != 0 {
      if this.Minutes != date.Minutes {
         return false
      }
   }

   return true
}

/// Checks if this date contains a 0 for month or day
func (this *Date) IsValid () bool {
   if this.Year == 0 || this.Month == 0 || this.Day == 0 {
      return false
   }
   return true
}

/// Checks if this date contains a 0 for all it's fields indicating
///   that this date as a property is to be considered empty and to be ignored.
func (this *Date) IsValidAsEmpty () bool {
   if this.Year == 0 && this.Month == 0 && this.Day == 0 && this.Hour == 0 && this.Minutes == 0 {
      return true
   }
   return false
}

/// Iterates the current day of this date by 1.
func (this *Date) IterateDay () {
   this.Day += 1

   if (this.Month == 1 || this.Month == 3 || this.Month == 5 ||
       this.Month == 7 || this.Month == 8 || this.Month == 10 ||
       this.Month == 12) && this.Day > 31 {
      this.Day = 1
      this.Month += 1
   }

   if (this.Month == 4 || this.Month == 6 || this.Month == 9 ||
       this.Month == 11) && this.Day > 30 {
      this.Day = 1
      this.Month += 1
   }

   if this.Month == 2 && this.Day > 28 {
      this.Day = 1
      this.Month += 1
   }

   if this.Month > 12 {
      this.Month = 1
      this.Year += 1
   }
}

/// Inequality check for <
func (this *Date) IsLessThan (date *Date) bool {

   if DebugLevel == 1 {
      fmt.Printf("Evaluating %d-%d-%d is less than %d-%d-%d",
         this.Year, this.Month, this.Day, date.Year, date.Month, date.Day)
   }

   if this.Year == date.Year && this.Month == date.Month && this.Day >= date.Day {
      return false
   }

   if this.Year == date.Year && this.Month > date.Month {
      return false
   }

   if this.Year > date.Year {
      return false
   }

   return true
}

/// Inequality check for <=
func (this *Date) IsLessThanOrEqualTo (date *Date) bool {

   if DebugLevel == 1 {
      fmt.Printf("Evaluating %d-%d-%d is less than or equal to %d-%d-%d",
         this.Year, this.Month, this.Day, date.Year, date.Month, date.Day)
   }

   if this.Year == date.Year && this.Month == date.Month && this.Day > date.Day {
      return false
   }

   if this.Year == date.Year && this.Month > date.Month {
      return false
   }

   if this.Year > date.Year {
      return false
   }

   return true
}

/// Inequality check for >
func (this *Date) IsGreaterThan (date *Date) bool {

   if DebugLevel == 1 {
      fmt.Printf("Evaluating %d-%d-%d is greater than %d-%d-%d",
         this.Year, this.Month, this.Day, date.Year, date.Month, date.Day)
   }

   if this.Year == date.Year && this.Month == date.Month && this.Day <= date.Day {
      return false
   }

   if this.Year == date.Year && this.Month < date.Month {
      return false
   }

   if this.Year < date.Year {
      return false
   }

   return true
}

/// Inequality check for >=
func (this *Date) IsGreaterThanOrEqualTo (date *Date) bool {

   if DebugLevel == 1 {
      log.Printf("Evaluating %d-%d-%d is greater than or equal to %d-%d-%d",
         this.Year, this.Month, this.Day, date.Year, date.Month, date.Day)
   }

   if this.Year == date.Year && this.Month == date.Month && this.Day < date.Day {
      if DebugLevel == 1 { log.Printf("Date is NOT greater than or equal - step 1") }
      return false
   }

   if this.Year == date.Year && this.Month < date.Month {
      if DebugLevel == 1 { log.Printf("Date is NOT greater than or equal - step 2") }
      return false
   }

   if this.Year < date.Year {
      if DebugLevel == 1 { log.Printf("Date is NOT greater than or equal - step 3") }
      return false
   }

   if DebugLevel == 1 { log.Printf("Date is greater and/or equal") }

   return true
}
