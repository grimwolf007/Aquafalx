#GUI
import tkinter as tk
#For the Time
from time import strftime
#For the sync Time
class Application(tk.Frame):
    def __init__(self, master=None):
        super().__init__(master)
        self.master = master
        self.pack()
        
        self.master.title('Ignisfalx')
        self.master.minsize(800,640) 

        self.create_widgets()
        self.refresh_time()
    def create_widgets(self):
        self.time = tk.Label(self)
        self.time["text"] = strftime('%H:%M:%S %p')
        self.time["region"] = "2 3 4 5"
        self.time.pack(side="top")

        self.quit = tk.Button(self, text="QUIT", fg="red",command=self.master.destroy, anchor="center")
        self.quit.pack(side="bottom")

    def refresh_time(self):
        self.time["text"] = strftime('%H:%M:%S %p')
        self.time.after(1000, self.refresh_time)
        

root = tk.Tk()
app = Application(master=root)


app.mainloop()

